package egg

import (
	"fmt"
	"net/http"
)

type Egg struct {
	debug  bool
	router *router
}

func (e *Egg) Debug() *Egg {
	e.debug = true
	return e
}

func (e *Egg) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(w, r)
	e.router.handle(ctx)
}

func New() *Egg {
	return &Egg{router: newRouter()}
}

func (e *Egg) Run(addr string) error {
	if e.debug {
		fmt.Println("[EGG] application run in debug mode.")
		e.router.print()
	}
	return http.ListenAndServe(addr, e)
}

func (e *Egg) RunTLS(addr, certFile, keyFile string) error {
	if e.debug {
		fmt.Println("[EGG] application run in debug mode.")
		e.router.print()
	}
	return http.ListenAndServeTLS(addr, certFile, keyFile, e)
}

func (e *Egg) GET(path string, handler HandleFunc) {
	e.router.addRoute("GET", path, handler)
}

func (e *Egg) POST(path string, handler HandleFunc) {
	e.router.addRoute("POST", path, handler)
}

func (e *Egg) PUT(path string, handler HandleFunc) {
	e.router.addRoute("PUT", path, handler)
}

func (e *Egg) DELETE(path string, handler HandleFunc) {
	e.router.addRoute("DELETE", path, handler)
}

func (e *Egg) PATCH(path string, handler HandleFunc) {
	e.router.addRoute("PATCH", path, handler)
}

func (e *Egg) OPTIONS(path string, handler HandleFunc) {
	e.router.addRoute("OPTIONS", path, handler)
}

func (e *Egg) HEAD(path string, handler HandleFunc) {
	e.router.addRoute("HEAD", path, handler)
}

func (e *Egg) Any(path string, handler HandleFunc) {
	e.GET(path, handler)
	e.POST(path, handler)
	e.PUT(path, handler)
	e.DELETE(path, handler)
	e.PATCH(path, handler)
	e.OPTIONS(path, handler)
	e.HEAD(path, handler)
}
