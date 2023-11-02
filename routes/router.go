package routes

import (
	"cc-backend-root-console/app/boot/route"
)

type Group struct {
	routes []route.Route
}

func (group *Group) AddRoute(route route.Route) {
	group.routes = append(group.routes, route)
}

func (group *Group) GetRoutes() []route.Route {
	return group.routes
}

func BuildGroupInstance() *Group {
	//loop through api array and add each route to the group
	group := &Group{}
	for _, r := range api {
		group.AddRoute(r)
	}
	return &Group{}
}
