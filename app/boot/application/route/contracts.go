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
	getHandler() *gin.HandlerFunc
}

type Group interface {
	GetRoutes() []Route
}

type Router interface {
	ApplyRoutes(Group)
}

func BuildRouterInstance(engine *gin.Engine) Router {
	return &routerInstance{engine: engine}
}

func BuildPathInstance(prefix string, relativePath string) Path {
	return &pathInstance{prefix: array.FilterStrings(strings.Split(prefix, "/"), func(s string) bool {
		return strings.Trim(s, " \n\t\r") != ""
	}), path: strings.Trim(relativePath, "/")}
}

func BuildRouteInstance(method Method, path Path, handler gin.HandlerFunc) Route {
	return &routeInstance{method: method, path: path, handler: &handler}
}
