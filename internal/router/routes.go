package router

import (
	"context"
	"log"
	"net/http"

	"github.com/mvndaai/ctxerr"
	"github.com/mvndaai/known-socially/internal/config"
	"github.com/mvndaai/known-socially/internal/db"
	"github.com/mvndaai/known-socially/internal/jwt"
	"github.com/mvndaai/known-socially/internal/router/server"
	"github.com/mvndaai/known-socially/internal/types"
	"github.com/mvndaai/validjson"
)

func StartServer() error {
	ctx := context.Background()

	db, err := db.New(ctx)
	if err != nil {
		return ctxerr.QuickWrap(ctx, err)
	}
	defer db.Close(ctx)

	h, err := NewHandler(ctx)
	if err != nil {
		return ctxerr.QuickWrap(ctx, err)
	}
	defer h.Close()

	rootRouter, err := server.New(
		server.Config[GenericHandlerFunc]{
			PathPrefix: "/",
			//GenericMiddleware     []func(T) T
			//Middleware            []MiddlewareFunc
			//DefaultParameters     openapi3.Parameters
			//AllowedOptionsHeaders []string
			GenericToHTTP: GenericToHTTP, // func(T) http.HandlerFunc
		},
		server.DocConfig{
			ServiceName: "known-socially",
			Description: "Link social media accounts",
			Version:     "0.0.1",
		})
	if err != nil {
		return ctxerr.Wrap(ctx, err, "120acdfb-98eb-4a65-a298-6619b8b7c942")
	}

	// Load the svelte static frontend files
	// https://codeandlife.com/2022/02/12/combine-golang-and-sveltekit-for-gui/
	rootRouter.Handle("", http.FileServer(http.Dir("./frontend/static")))

	rootRouter.Endpoint("/status", http.MethodGet, statusHandler, nil)
	apiRouter := rootRouter.Subrouter(server.Config[GenericHandlerFunc]{
		PathPrefix: "/api",
	})
	apiRouter.Endpoint("/domain", http.MethodGet, h.domainListHandler, nil)
	//apiRouter.Endpoint("/domain/{id}", http.MethodGet, h.domainGetHandler, nil)
	apiRouter.Endpoint("/user", http.MethodGet, h.userListHandler, nil)

	jwtMiddleware := JWTMiddleware(db)
	protectedapiRouter := apiRouter.Subrouter(server.Config[GenericHandlerFunc]{
		PathPrefix: "/protected",
		Middleware: []server.MiddlewareFunc{jwtMiddleware, JWTSubjectMiddleware},
	})
	protectedapiRouter.Endpoint("/domain", http.MethodPost, h.domainCreateHandler, nil)
	protectedapiRouter.Endpoint("/user", http.MethodPost, h.userCreateHandler, nil)
	protectedapiRouter.Endpoint("/logout", http.MethodPost, h.logoutHandler, nil)

	env := config.GetEnviroment()
	if env == "dev" {
		testRouter := apiRouter.Subrouter(server.Config[GenericHandlerFunc]{
			PathPrefix: "/test",
		})
		testRouter.Endpoint("/error", http.MethodGet, testErrorHandler, nil)
		testRouter.Endpoint("/error", http.MethodPost, testErrorHandler, nil)
		testRouter.Endpoint("/jwt", http.MethodPost, testCreateJWTHandler, nil)
		testRouter.Endpoint("/list", http.MethodGet, func(r *http.Request) (data, meta any, status int, _ error) {
			return rootRouter.ListRoutes(), nil, http.StatusOK, nil
		}, nil)

		apiRouter.Subrouter(server.Config[GenericHandlerFunc]{
			PathPrefix: "/test/auth",
			Middleware: []server.MiddlewareFunc{jwtMiddleware},
		}).Endpoint("", http.MethodGet, statusHandler, nil)

	}

	port := config.GetPort()
	s := rootRouter.NewServer(port, nil)
	log.Printf("Starting '%s' server at http://localhost%s\n", env, port)
	log.Fatal(s.ListenAndServe())
	return nil
}

func statusHandler(r *http.Request) (data, meta any, status int, _ error) {
	return nil, nil, http.StatusOK, nil
}

func testErrorHandler(r *http.Request) (data, meta any, status int, _ error) {
	return nil, nil, http.StatusBadGateway, ctxerr.New(r.Context(), "72c7374f-4ba6-41db-acad-1741913422dd", "test error")
}

func testCreateJWTHandler(r *http.Request) (data, meta any, status int, _ error) {
	ctx := r.Context()
	claims := jwt.JWTClaims{}
	err := validjson.UnmarshalReadCloser(ctx, r.Body, &claims)
	defer r.Body.Close()
	if err != nil {
		return nil, nil, http.StatusBadRequest, ctxerr.Wrap(ctx, err, "b0851ad6-83aa-4bae-a492-7473e527fe2f")
	}

	token, err := jwt.GenerateJWT(ctx, claims)
	if err != nil {
		return nil, nil, http.StatusBadGateway, ctxerr.Wrap(ctx, err, "8219cb01-e2cb-41d9-8d0a-6512faa04441")
	}
	return token, nil, http.StatusOK, nil
}

func (h *Handler) domainCreateHandler(r *http.Request) (data, meta any, status int, _ error) {
	ctx := r.Context()
	body := types.DomainCreate{}
	err := validjson.UnmarshalReadCloser(ctx, r.Body, &body)
	defer r.Body.Close()
	if err != nil {
		return nil, nil, http.StatusBadRequest, ctxerr.Wrap(ctx, err, "9ac2a9a9-a580-403d-8d27-c956664ea39b")
	}

	d, err := h.db.CreateDomain(ctx, body)
	if err != nil {
		return nil, nil, http.StatusBadRequest, ctxerr.QuickWrap(ctx, err)
	}

	return d, nil, http.StatusOK, nil
}

func (h *Handler) userCreateHandler(r *http.Request) (data, meta any, status int, _ error) {
	ctx := r.Context()
	body := types.UserCreate{}
	err := validjson.UnmarshalReadCloser(ctx, r.Body, &body)
	defer r.Body.Close()
	if err != nil {
		return nil, nil, http.StatusBadRequest, ctxerr.Wrap(ctx, err, "36243519-91ef-470f-81e5-13173d29be87")
	}

	d, err := h.db.CreateUser(ctx, body)
	if err != nil {
		return nil, nil, http.StatusBadRequest, ctxerr.QuickWrap(ctx, err)
	}

	return d, nil, http.StatusOK, nil
}

func (h *Handler) domainListHandler(r *http.Request) (data, meta any, status int, _ error) {
	ctx := r.Context()
	list := types.DomainList{}
	err := list.Fill(ctx, r.URL.Query())
	domains, pagination, err := h.db.ListDomains(ctx, list.Filters, list.Pagination)
	if err != nil {
		return nil, nil, http.StatusBadRequest, ctxerr.QuickWrap(ctx, err)
	}
	return domains, pagination, http.StatusOK, nil
}

func (h *Handler) userListHandler(r *http.Request) (data, meta any, status int, _ error) {
	ctx := r.Context()
	list := types.UserList{}
	err := list.Fill(ctx, r.URL.Query())
	users, pagination, err := h.db.ListUsers(ctx, list.Filters, list.Pagination)
	if err != nil {
		return nil, nil, http.StatusBadRequest, ctxerr.QuickWrap(ctx, err)
	}
	return users, pagination, http.StatusOK, nil
}
