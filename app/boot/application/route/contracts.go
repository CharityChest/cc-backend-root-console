package route

import (
	"cc-backend-root-console/app/helpers/array"
	"github.com/gin-gonic/gin"
	"strings"
)

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
	PATCH  Method = "PATCH"
)

type Path interface {
	getPrefix() []string
	getRelativePath() string
	getFullPath() string
	isEqual(Path) bool
}

type Route interface {
	getMethod() Method
	getPath() Path
	getHandler() *func(*gin.Context)
}

type Group interface {
	GetRoutes() []Route
}

type Router interface {
	ApplyRoutes(Group)
}

func BuildRouterInstance(engine *gin.Engine) Router {
	return &RouterInstance{engine: engine}
}

func BuildPathInstance(prefix string, relativePath string) Path {
	return &PathInstance{prefix: array.FilterStrings(strings.Split(prefix, "/"), func(s string) bool {
		return strings.Trim(s, " \n\t\r") != ""
	}), path: strings.Trim(relativePath, "/")}
}

func BuildRouteInstance(method Method, path Path, handler func(*gin.Context)) Route {
	return &RouteInstance{method: method, path: path, handler: &handler}
}
