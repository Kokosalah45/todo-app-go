package types

import "net/http"

type  RouteRegisterer interface {
	RegisterRoutes(router *http.ServeMux)
}

