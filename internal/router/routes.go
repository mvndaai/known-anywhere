package router

import (
	"log"
	"net/http"

	"github.com/mvndaai/ctxerr"
)

func StartServer() {
	NewRoute(http.MethodGet, "/", ListRoutesHandler)
	NewRoute(http.MethodGet, "/status", statusHandler)
	NewRoute(http.MethodGet, "/test/error", testErrorHandler)

	port := GetPort()
	log.Println("Starting server at http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func statusHandler(r *http.Request) (data, meta any, status int, _ error) {
	return nil, nil, http.StatusOK, nil
}

func testErrorHandler(r *http.Request) (data, meta any, status int, _ error) {
	return nil, nil, http.StatusBadGateway, ctxerr.New(r.Context(), "c691a249-1413-407d-a309-b44d52eb51ac", "test error")
}
