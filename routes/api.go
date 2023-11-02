package routes

import (
	"cc-backend-root-console/app/boot/route"
	"cc-backend-root-console/app/controllers/organization"
)

var organizationController = organization.BuildOrganizationController()
var organizationControllerListMethod = organizationController.List

var api = []route.Route{
	route.BuildRouteInstance(
		route.GET,
		route.BuildPathInstance("/api/v1/", "organizations"),
		&organizationControllerListMethod,
	),
}
