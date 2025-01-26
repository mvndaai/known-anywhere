package router

import (
	"log"
	"net/http"

	"github.com/mvndaai/ctxerr"
	"github.com/mvndaai/known-socially/internal/jwt"
	"github.com/mvndaai/validjson"
)

func StartServer() {

	// Load html files
	// https://codeandlife.com/2022/02/12/combine-golang-and-sveltekit-for-gui/
	http.Handle("/", http.FileServer(http.Dir("./frontend/static")))

	NewRoute(http.MethodGet, "/api", ListRoutesHandler)
	NewRoute(http.MethodGet, "/status", statusHandler)
	NewRoute(http.MethodGet, "/test/error", testErrorHandler)
	NewRoute(http.MethodPost, "/test/jwt", testCreateJWTHandler)
	NewRoute(http.MethodGet, "/test/auth", statusHandler, JWTMiddleware)

	port := GetPort()
	log.Println("Starting server at http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
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
