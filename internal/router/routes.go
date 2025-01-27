package router

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/mvndaai/ctxerr"
	"github.com/mvndaai/known-socially/internal/jwt"
	"github.com/mvndaai/known-socially/internal/types"
	"github.com/mvndaai/validjson"
)

func StartServer() error {
	// Load html files
	// https://codeandlife.com/2022/02/12/combine-golang-and-sveltekit-for-gui/
	http.Handle("/", http.FileServer(http.Dir("./frontend/static")))

	h, err := NewHandler()
	if err != nil {
		return ctxerr.QuickWrap(context.Background(), err)
	}
	defer h.Close()

	NewRoute(http.MethodGet, "/status", statusHandler)
	NewRoute(http.MethodGet, "/api", ListRoutesHandler)

	env := os.Getenv("ENVIRONMENT")
	if env == "dev" {
		// TODO make a better mux
		NewRoute(http.MethodGet, "/test/error", testErrorHandler)
		NewRoute(http.MethodPost, "/test/jwt", testCreateJWTHandler)
		NewRoute(http.MethodGet, "/test/auth", statusHandler, JWTMiddleware)

		// Most get/list routes are unprotected
		NewRoute(http.MethodGet, "/api/domain", h.domainListHandler)
		//NewRoute(http.MethodGet, "/api/domain/{id}", h.domainGetHandler)

		NewRoute(http.MethodPost, "/api/protected/domain", h.domainCreateHandler, JWTMiddleware)
	}

	port := GetPort()
	log.Println("Starting server at http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
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
		return nil, nil, http.StatusBadGateway, ctxerr.Wrap(ctx, err, "ea32c534-2c39-4a92-9714-0863073a79c6")
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

func (h *Handler) domainListHandler(r *http.Request) (data, meta any, status int, _ error) {
	ctx := r.Context()
	list := types.DomainList{}
	err := list.Fill(ctx, r.URL.Query())
	domains, pagination, err := h.db.ListDomains(ctx, list)
	if err != nil {
		return nil, nil, http.StatusBadRequest, ctxerr.QuickWrap(ctx, err)
	}
	return domains, pagination, http.StatusOK, nil
}
