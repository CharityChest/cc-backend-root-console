package route

import "github.com/gin-gonic/gin"

type routerInstance struct {
	engine *gin.Engine
}

func (r *routerInstance) ApplyRoutes(group Group) {
	routes := group.GetRoutes()
	for _, route := range routes {
		method := route.getMethod()
		path := route.getPath().getFullPath()
		handler := route.getHandler()
		r.engine.Handle(string(method), path, *handler)
	}
}
