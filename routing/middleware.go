package routing

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	h "nyded/helpers"
)

type ValidateRouteMiddleware struct {
	next   http.Handler
	router *httprouter.Router
}

func newValidateRouteMiddleware(next http.Handler, router *httprouter.Router) *ValidateRouteMiddleware {
	return &ValidateRouteMiddleware{next: next, router: router}
}

func (m *ValidateRouteMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	routeValid, _, _ := m.router.Lookup(r.Method, r.URL.Path)
	if routeValid != nil {
		m.next.ServeHTTP(w, r)
		return
	}
	h.Respond(w, nil, http.StatusNotFound)
}
