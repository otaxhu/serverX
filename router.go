package serverX

import "net/http"

type Router struct {
	Rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		Rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHandler(path, method string) (handler http.HandlerFunc, methodExist bool, pathExist bool) {
	_, pathExist = r.Rules[path]
	handler, methodExist = r.Rules[path][method]
	return
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, methodExist, pathExist := r.FindHandler(request.URL.Path, request.Method)
	if !pathExist {
		w.WriteHeader(http.StatusNotFound)
	} else if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		handler(w, request)
	}

}
