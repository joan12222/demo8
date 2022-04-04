package srv

import (
	"net/http"
)

type Routable interface {
	Route(method string, path string, handler func(c *Context))
}

type Server interface {
	Routable
	Start(address string)
}

type sdkHttpServer struct {
	name string
	maph *HandlersBasedOnMap
}

// Route: add route
func (s *sdkHttpServer) Route(method string, path string, handler handlerFunc) {
	key := s.maph.key(method, path)
	s.maph.handlers.Store(key, handler)
}

func (s *sdkHttpServer) Start(address string) {
	// as long as s.maph implements method "ServeHTTP", it can be used by http.Handle
	http.Handle("/", s.maph)
	// nil means it uses DefautServerMux
	http.ListenAndServe(address, nil)
}

func NewHttpServer(n string) *sdkHttpServer {
	return &sdkHttpServer{
		name: n,
		maph: NewHandlersBasedOnMap(),
	}
}
