package server_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"testing"

	"github.com/mvndaai/known-socially/internal/router/server"
	"github.com/stretchr/testify/assert"
)

type GenericHandlerFunc func(r *http.Request) (data, meta any, status int, _ error)

func GenericToHTTP(ghf GenericHandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		data, meta, status, err := ghf(req)
		if err != nil {
			http.Error(res, err.Error(), status)
			return
		}
		resp := map[string]any{
			"data": data,
			"meta": meta,
		}
		res.Header().Set("Content-Type", "application/json")
		b, err := json.Marshal(resp)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = res.Write(b)
		if status == 0 {
			res.WriteHeader(status)
		}
	}
}
func GenericHandle(ghf GenericHandlerFunc) GenericHandlerFunc {
	return ghf
}

func TestGenericRouter(t *testing.T) {
	server.New(server.Config[GenericHandlerFunc]{
		PathPrefix: "/root",
	}, server.DocConfig{})
}

func findOpenPort(t *testing.T) string {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	port := fmt.Sprintf(":%d", listener.Addr().(*net.TCPAddr).Port)
	err = listener.Close()
	if err != nil {
		t.Fatal(err)
	}
	return port
}

func TestRouter(t *testing.T) {
	// Create middleware
	var calledRootMiddleware bool
	rootMiddleware := func(l http.HandlerFunc) http.HandlerFunc {
		return func(res http.ResponseWriter, req *http.Request) {
			t.Log("in root middleware")
			calledRootMiddleware = true
			l(res, req)
		}
	}
	var calledSubMiddleware bool
	var calledRootMilddlewareBeforeSub bool
	subMiddleware := func(l http.HandlerFunc) http.HandlerFunc {
		return func(res http.ResponseWriter, req *http.Request) {
			t.Log("in sub middleware")
			if calledRootMiddleware {
				calledRootMilddlewareBeforeSub = true
			}
			calledSubMiddleware = true
			l(res, req)
		}
	}

	// Create router and subrouters
	basePath := "/root"
	rr, err := server.New(server.Config[GenericHandlerFunc]{
		PathPrefix:    basePath,
		Middleware:    []server.MiddlewareFunc{rootMiddleware},
		GenericToHTTP: GenericToHTTP,
		//GenericHandle: func(ghf GenericHandlerFunc) GenericHandlerFunc {},
	}, server.DocConfig{
		ServiceName: "test",
		Description: "test",
		Version:     "v0.0.1",
	})
	if err != nil {
		t.Fatal(err)
	}
	sub1Path := "/sub1"
	sub1 := rr.Subrouter(server.Config[GenericHandlerFunc]{
		PathPrefix: sub1Path,
		Middleware: []server.MiddlewareFunc{subMiddleware},
	})
	sub2Path := "/sub2"
	sub2 := sub1.Subrouter(server.Config[GenericHandlerFunc]{
		PathPrefix:            sub2Path,
		AllowedOptionsHeaders: []string{"X-Test-Header"},
	})

	// Create endpoint
	var calledEndpoint bool
	endpointPath := "/endpoint"
	laend := func(r *http.Request) (data, meta any, status int, _ error) {
		t.Log("in endpoint")
		calledEndpoint = true
		return
	}

	sub2.Endpoint(endpointPath, http.MethodGet, laend, nil)
	sub2.Endpoint(endpointPath, http.MethodPost, laend, nil)

	// Get an open port on the machine
	port := findOpenPort(t)
	t.Log("port", port)

	// Start server and stop server in gofuncs
	stopServer := make(chan struct{})
	blockUntilServerStopped := make(chan struct{})
	server := rr.NewServer(port, nil)
	go func() {
		log.Println("starting service")
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()
	go func() {
		<-stopServer
		log.Println("stopping service")
		if err := server.Shutdown(context.Background()); err != nil {
			panic(err)
		}
		close(blockUntilServerStopped)
	}()

	// Make http call to enpoint
	path := fmt.Sprintf("http://localhost%s%s%s%s%s", port, basePath, sub1Path, sub2Path, endpointPath)
	t.Log("path", path)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, path, http.NoBody)
	assert.Nil(t, err)
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	err = resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, []string{"GET,OPTIONS,POST"}, resp.Header["Access-Control-Allow-Methods"])
	assert.Contains(t, resp.Header["Access-Control-Allow-Headers"][0], ",X-Test-Header")

	req, err = http.NewRequestWithContext(context.Background(), http.MethodOptions, path, http.NoBody)
	assert.Nil(t, err)
	resp, err = http.DefaultClient.Do(req)
	assert.Nil(t, err)
	err = resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, []string{"GET,OPTIONS,POST"}, resp.Header["Access-Control-Allow-Methods"])
	assert.Contains(t, resp.Header["Access-Control-Allow-Headers"][0], ",X-Test-Header")

	// Stop server
	close(stopServer)
	<-blockUntilServerStopped

	// Check assertions
	assert.True(t, calledRootMiddleware)
	assert.True(t, calledSubMiddleware)
	assert.True(t, calledRootMilddlewareBeforeSub)
	assert.True(t, calledEndpoint)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

//func TestHealthCheck(t *testing.T) {
//	// Create router
//	rr, err := server.New(server.Config{
//		PathPrefix: "/root",
//	}, server.DocConfig{
//		ServiceName: "test",
//		Description: "test",
//		Version:     "v0.0.1",
//	})
//	if err != nil {
//		t.Fatal(err)
//	}

//	// Get an open port on the machine
//	port := findOpenPort(t)
//	t.Log("port", port)

//	// Start server and stop server in gofuncs
//	stopServer := make(chan struct{})
//	blockUntilServerStopped := make(chan struct{})
//	server := rr.NewServer(port, nil)
//	go func() {
//		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
//			panic(err)
//		}
//	}()
//	go func() {
//		<-stopServer
//		if err := server.Shutdown(context.Background()); err != nil {
//			panic(err)
//		}
//		close(blockUntilServerStopped)
//	}()

//	// Make http call to enpoint
//	path := fmt.Sprintf("http://localhost%s/healthcheck", port)
//	t.Log("path", path)
//	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, path, http.NoBody)
//	assert.Nil(t, err)
//	resp, err := http.DefaultClient.Do(req)
//	assert.Nil(t, err)
//	err = resp.Body.Close()
//	assert.Nil(t, err)
//	assert.Equal(t, http.StatusOK, resp.StatusCode)

//	// Stop server
//	close(stopServer)
//	<-blockUntilServerStopped
//}

//func TestDefaultMiddleware(t *testing.T) {
//	// Create router
//	basePath := "/root"
//	rr, err := server.New(server.Config{
//		PathPrefix: basePath,
//	}, server.DocConfig{
//		ServiceName: "test",
//		Description: "test",
//		Version:     "v0.0.1",
//	})
//	if err != nil {
//		t.Fatal(err)
//	}

//	keyQueryParam := "query"
//	//keyPathParam := "path"

//	var actualPathParam string
//	var actualQueryParam string

//	//end := func(req gsrequest.Request) (_ gsresponse.Envelope, _ error) {
//	//	t.Log("in endpoint")
//	//	q := req.MultiValueParameters
//	//	actualPathParam = q.Get(keyPathParam)
//	//	actualQueryParam = q.Get(keyQueryParam)

//	//	// Test LogHeadersAndParams
//	//	assert.Equal(t, "a", ctxerr.Fields(req.Context)["parameters"].(url.Values).Get(keyPathParam))
//	//	return
//	//}

//	//endpointPath := fmt.Sprintf("/{%s}", keyPathParam)
//	//rr.GSEndpoint(endpointPath, http.MethodGet, end, nil)

//	// Get an open port on the machine
//	port := findOpenPort(t)
//	t.Log("port", port)

//	// Start server and stop server in gofuncs
//	stopServer := make(chan struct{})
//	blockUntilServerStopped := make(chan struct{})
//	server := rr.NewServer(port, nil)
//	go func() {
//		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
//			panic(err)
//		}
//	}()
//	go func() {
//		<-stopServer
//		if err := server.Shutdown(context.Background()); err != nil {
//			panic(err)
//		}
//		close(blockUntilServerStopped)
//	}()

//	inputPathParam := " a "
//	inputQueryParam := " b "

//	// Make http call to enpoint
//	path := fmt.Sprintf("http://localhost%s%s/%s", port, basePath, inputPathParam)
//	t.Log("path", path)
//	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, path, http.NoBody)
//	q := req.URL.Query()
//	q.Add(keyQueryParam, inputQueryParam)
//	req.URL.RawQuery = q.Encode()

//	assert.Nil(t, err)
//	resp, err := http.DefaultClient.Do(req)
//	assert.Nil(t, err)
//	err = resp.Body.Close()
//	assert.Nil(t, err)

//	// Stop server
//	close(stopServer)
//	<-blockUntilServerStopped

//	// assertions
//	assert.Equal(t, http.StatusOK, resp.StatusCode)
//	assert.Equal(t, "a", actualPathParam)
//	assert.Equal(t, "b", actualQueryParam)
//}

//func TestRouterErrors(t *testing.T) {
//	_, err := server.New(server.Config{
//		PathPrefix: "/root",
//	}, server.DocConfig{
//		ServiceName: "",
//		Description: "test",
//		Version:     "v0.0.1",
//	})
//	assert.ErrorContains(t, err, "service name")

//	_, err = server.New(server.Config{
//		PathPrefix: "/root",
//	}, server.DocConfig{
//		ServiceName: "test",
//		Description: "",
//		Version:     "v0.0.1",
//	})
//	assert.ErrorContains(t, err, "description")

//	_, err = server.New(server.Config{
//		PathPrefix: "/root",
//	}, server.DocConfig{
//		ServiceName: "test",
//		Description: "test",
//		Version:     "",
//	})
//	assert.ErrorContains(t, err, "version")
//}

//func TestNoPathPrefix(t *testing.T) {
//	// Create router
//	basePath := ""

//	rr, err := server.New(server.Config{
//		PathPrefix:        basePath,
//		DefaultParameters: openapi3.Parameters{},
//	}, server.DocConfig{
//		ServiceName: "sn",
//		Description: "service description",
//		Version:     "v0.0.1",
//	})
//	if err != nil {
//		t.Fatal(err)
//	}

//	//end := func(_ gsrequest.Request) (_ gsresponse.Envelope, _ error) { return }

//	//enpointPath := "/endpoint"
//	//rr.GSEndpoint(enpointPath, http.MethodGet, end, nil)

//	// Get an open port on the machine
//	port := findOpenPort(t)
//	t.Log("port", port)

//	// Start server and stop server in gofuncs
//	stopServer := make(chan struct{})
//	blockUntilServerStopped := make(chan struct{})
//	server := rr.NewServer(port, nil)
//	go func() {
//		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
//			panic(err)
//		}
//	}()
//	go func() {
//		<-stopServer
//		if err := server.Shutdown(context.Background()); err != nil {
//			panic(err)
//		}
//		close(blockUntilServerStopped)
//	}()

//	// Make http call to swagger enpoint
//	path := fmt.Sprintf("http://localhost%s/endpoint", port)
//	t.Log("path", path)
//	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, path, http.NoBody)
//	assert.Nil(t, err)
//	resp, err := http.DefaultClient.Do(req)
//	assert.Nil(t, err)
//	err = resp.Body.Close()
//	assert.Nil(t, err)

//	// Stop server
//	close(stopServer)
//	<-blockUntilServerStopped

//	// assertions
//	assert.Equal(t, http.StatusOK, resp.StatusCode)
//}
