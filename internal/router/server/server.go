package server

import (
	"fmt"
	"net/http"
	"path"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
)

var defaultMiddleware = []MiddlewareFunc{
	TrimSpaces,
	LogHeadersAndParams,
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc
type DocFunc func() (*openapi3.Operation, error)

type Router[T any] interface {
	Subrouter(Config[T]) Router[T]
	Endpoint(endpointPath, method string, handler T, doc DocFunc)
	Handle(path string, handler http.Handler)
}

type RootRouter[T any] interface {
	Router[T]
	NewServer(port string, sc *ServerConfig) *http.Server
	ListRoutes() map[string][]string
}

type router[T any] struct {
	routeMux              *routeMux
	doc                   *openapi3.T
	pathPrefix            string
	statsHandler          func(next http.HandlerFunc, method, fullPath string) http.HandlerFunc
	genericMiddleware     []func(T) T
	middleware            []MiddlewareFunc
	defaultParameters     openapi3.Parameters
	allowedOptionsHeaders []string
	genericToHTTP         func(T) http.HandlerFunc
}

type rootrouter[T any] struct {
	router[T]
}

type Config[T any] struct {
	PathPrefix            string
	GenericMiddleware     []func(T) T
	Middleware            []MiddlewareFunc
	DefaultParameters     openapi3.Parameters
	AllowedOptionsHeaders []string
	GenericToHTTP         func(T) http.HandlerFunc
}

type DocConfig struct {
	ServiceName string
	Description string
	Version     string
	Tags        []*openapi3.Tag
}

func (dc DocConfig) validate() error {
	if dc.ServiceName == "" {
		return fmt.Errorf("service name required")
	}
	if dc.Description == "" {
		return fmt.Errorf("description required")
	}
	if dc.Version == "" {
		return fmt.Errorf("version required")
	}
	return nil
}

type ServerConfig struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func New[T any](rc Config[T], dc DocConfig) (RootRouter[T], error) {
	if err := dc.validate(); err != nil {
		return nil, fmt.Errorf("doc config: %w", err)
	}

	docBase := &openapi3.T{
		OpenAPI: "3.0.3",
		Info: &openapi3.Info{
			Title:       dc.ServiceName,
			Description: dc.Description,
			Version:     dc.Version,
		},
		Tags: dc.Tags,
	}

	if rc.GenericToHTTP == nil {
		return nil, fmt.Errorf("missing GenericToHTTP in config")
	}

	if rc.PathPrefix != "" {
		return &rootrouter[T]{
			router: router[T]{
				routeMux:              newRouteMux(),
				doc:                   docBase,
				pathPrefix:            rc.PathPrefix,
				middleware:            append(defaultMiddleware, rc.Middleware...),
				genericMiddleware:     rc.GenericMiddleware,
				defaultParameters:     rc.DefaultParameters,
				allowedOptionsHeaders: rc.AllowedOptionsHeaders,
				genericToHTTP:         rc.GenericToHTTP,
			},
		}, nil
	}
	return nil, fmt.Errorf("path prefix required")
}

func (r *router[T]) Subrouter(rc Config[T]) Router[T] {
	subRouter := &router[T]{
		routeMux:              r.routeMux, // Share the same routeMux
		doc:                   r.doc,
		pathPrefix:            path.Join(r.pathPrefix, rc.PathPrefix),
		middleware:            append(r.middleware, rc.Middleware...),
		defaultParameters:     append(r.defaultParameters, rc.DefaultParameters...),
		allowedOptionsHeaders: append(r.allowedOptionsHeaders, rc.AllowedOptionsHeaders...),
		genericMiddleware:     append(r.genericMiddleware, rc.GenericMiddleware...),
		genericToHTTP:         r.genericToHTTP,
	}

	if subRouter.genericToHTTP == nil {
		subRouter.genericToHTTP = r.genericToHTTP
	}

	return subRouter
}

func (r *router[T]) Endpoint(endpointPath, method string, handler T, doc DocFunc) {
	fullPath := path.Join(r.pathPrefix, endpointPath)

	// Apply generic middleware in reverse order
	for i := len(r.genericMiddleware) - 1; i >= 0; i-- {
		handler = r.genericMiddleware[i](handler)
	}

	//// Apply default generic handler
	//handler = genericHandle(handler)

	// Convert to HTTP handler
	httpHandler := r.genericToHTTP(handler)

	// Apply HTTP middleware in reverse order
	for i := len(r.middleware) - 1; i >= 0; i-- {
		httpHandler = r.middleware[i](httpHandler)
	}

	if r.statsHandler != nil {
		httpHandler = r.statsHandler(httpHandler, method, fullPath)
	}

	// Replace route addition with routeMux
	r.routeMux.addHandler(fullPath, method, httpHandler)

	// Add swagger docs
	if doc != nil {
		op, err := doc()
		if err != nil {
			panic(err)
		}
		op.Parameters = append(op.Parameters, r.defaultParameters...)
		addDocPath(fullPath, method, r.doc, op)
	}
}

// Add new method to router type
func (r *router[T]) Handle(path string, handler http.Handler) {
	fullPath := r.pathPrefix + path
	r.routeMux.addHandler(fullPath, "", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}))
}

func (rr *rootrouter[T]) NewServer(port string, sc *ServerConfig) *http.Server {
	mux := http.NewServeMux()

	paths := rr.routeMux.getAllPaths()
	for _, path := range paths {
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			handler, exists := rr.routeMux.getHandler(path, r.Method)
			if !exists {
				// For non-method specific handlers (like FileServer)
				handler, exists = rr.routeMux.getHandler(path, "")
				if !exists {
					if r.Method != http.MethodOptions {
						w.WriteHeader(http.StatusMethodNotAllowed)
						return
					}
					methods := rr.routeMux.getMethods(path)
					handler = optionsHandler(methods, rr.allowedOptionsHeaders)
				}
			}
			handler(w, r)
		})
	}

	if sc == nil {
		sc = &ServerConfig{}
	}
	if sc.ReadTimeout == 0 {
		sc.ReadTimeout = 5 * time.Second
	}
	if sc.WriteTimeout == 0 {
		sc.WriteTimeout = 90 * time.Second
	}
	if sc.IdleTimeout == 0 {
		sc.IdleTimeout = 10 * time.Minute
	}

	return &http.Server{
		Handler:      mux,
		ReadTimeout:  sc.ReadTimeout,
		WriteTimeout: sc.WriteTimeout,
		IdleTimeout:  sc.IdleTimeout,
		Addr:         port,
	}
}

func (rr *rootrouter[T]) ListRoutes() map[string][]string {
	allPaths := rr.routeMux.getAllPaths()
	ret := make(map[string][]string)
	for _, path := range allPaths {
		methods := rr.routeMux.getMethods(path)
		ret[path] = methods
	}
	return ret
}

func addDocPath(path, method string, s *openapi3.T, op *openapi3.Operation) {
	p := s.Paths.Value(path)
	if p == nil {
		p = &openapi3.PathItem{}
	}

	switch method {
	case http.MethodGet:
		p.Get = op
	case http.MethodPost:
		p.Post = op
	case http.MethodPut:
		p.Put = op
	case http.MethodPatch:
		p.Patch = op
	case http.MethodDelete:
		p.Delete = op
	}

	s.Paths.Set(path, p)
}
