package route

import "github.com/gin-gonic/gin"

type RouteInstance struct {
	method  Method
	path    Path
	handler *func(*gin.Context)
}

func (r *RouteInstance) getMethod() Method {
	return r.method
}

func (r *RouteInstance) getPath() Path {
	return r.path
}

func (r *RouteInstance) getHandler() *func(*gin.Context) {
	return r.handler
}
