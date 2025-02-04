package server

import (
	"net/http"
	"sync"
)

type methodHandler struct {
	handlers map[string]http.HandlerFunc
	mu       sync.RWMutex
}

// New type to encapsulate route handling
type routeMux struct {
	routes map[string]*methodHandler
	mu     sync.RWMutex
}

func newRouteMux() *routeMux {
	return &routeMux{
		routes: make(map[string]*methodHandler),
	}
}

func (rm *routeMux) getHandler(path, method string) (http.HandlerFunc, bool) {
	rm.mu.RLock()
	defer rm.mu.RUnlock()

	if mh, exists := rm.routes[path]; exists {
		mh.mu.RLock()
		defer mh.mu.RUnlock()
		h, exists := mh.handlers[method]
		return h, exists
	}
	return nil, false
}

func (rm *routeMux) addHandler(path, method string, handler http.HandlerFunc) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if rm.routes[path] == nil {
		rm.routes[path] = &methodHandler{
			handlers: make(map[string]http.HandlerFunc),
		}
	}

	rm.routes[path].mu.Lock()
	rm.routes[path].handlers[method] = handler
	rm.routes[path].mu.Unlock()
}

func (rm *routeMux) getMethods(path string) []string {
	rm.mu.RLock()
	defer rm.mu.RUnlock()

	if mh, exists := rm.routes[path]; exists {
		mh.mu.RLock()
		defer mh.mu.RUnlock()

		methods := make([]string, 0, len(mh.handlers))
		for method := range mh.handlers {
			methods = append(methods, method)
		}
		return methods
	}
	return nil
}

func (rm *routeMux) getAllPaths() []string {
	rm.mu.RLock()
	defer rm.mu.RUnlock()

	paths := make([]string, 0, len(rm.routes))
	for path := range rm.routes {
		paths = append(paths, path)
	}
	return paths
}
