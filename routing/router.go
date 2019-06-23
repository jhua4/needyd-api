package routing

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"net/http"
)

// NewRouter Create router and assign routes
func NewRouter() http.Handler {
	router := httprouter.New()
	router.GET("/", index)

	route := newValidateRouteMiddleware(router, router)
	cors := cors.AllowAll().Handler(route)

	return cors
}
