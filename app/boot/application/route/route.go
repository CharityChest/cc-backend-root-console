package route

import "github.com/gin-gonic/gin"

type routeInstance struct {
	method  Method
	path    Path
	handler *gin.HandlerFunc
}

func (r *routeInstance) getMethod() Method {
	return r.method
}

func (r *routeInstance) getPath() Path {
	return r.path
}

func (r *routeInstance) getHandler() *gin.HandlerFunc {
	return r.handler
}
