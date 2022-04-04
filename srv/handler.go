package srv

import (
	"fmt"
	"net/http"
	"sync"
)

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

var _ Handler = &HandlersBasedOnMap{}

type handlerFunc func(c *Context)

type HandlersBasedOnMap struct {
	handlers sync.Map
}

// run handler
func (h *HandlersBasedOnMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r)
	key := h.key(r.Method, r.URL.Path)
	handler, ok := h.handlers.Load(key)
	if !ok {
		c.W.WriteHeader(http.StatusNotFound)
		_, _ = c.W.Write([]byte("no route for this key"))
		return
	}
	handler.(handlerFunc)(c)
}

func (h *HandlersBasedOnMap) key(method string, path string) string {
	return fmt.Sprintf("#{method}##{path}")
}

func NewHandlersBasedOnMap() *HandlersBasedOnMap {
	return &HandlersBasedOnMap{}
}
