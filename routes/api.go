package routes

import (
	"cc-backend-root-console/app/boot/application/route"
	"cc-backend-root-console/app/controllers/organization"
)

const API_V1_PREFIX = "/api/v1/"

var organizationController = organization.BuildOrganizationController()
var organizationControllerListMethod = organizationController.List

var api = []route.Route{
	route.BuildRouteInstance(
		route.GET,
		route.BuildPathInstance(API_V1_PREFIX, "organizations"),
		&organizationControllerListMethod,
	),
}
