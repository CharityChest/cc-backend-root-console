package routes

import (
	"cc-backend-root-console/app/boot/application/route"
	"cc-backend-root-console/app/http/controllers/organization"
)

const ApiV1Prefix = "/api/v1/"

var organizationController = organization.BuildOrganizationController()

var api = []route.Route{
	route.BuildRouteInstance(
		route.GET,
		route.BuildPathInstance(ApiV1Prefix, "organizations"),
		organizationController.List,
	),
}
