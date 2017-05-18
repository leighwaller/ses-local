package route

import (
	"github.com/gorilla/mux"
	"net/http"
	"ses-local/handler"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, r := range routes {
		h := handler.RequestLoggingAdapter(r.handler, r.name)
		router.Methods(r.method).Path(r.path).Name(r.name).Handler(h)
	}
	return router
}

type Route struct {
	name    string
	method  string
	path    string
	handler http.HandlerFunc
}

var routes = []Route{
	{"index", "POST", "/", handler.Index},
}
