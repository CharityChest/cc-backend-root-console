package organization

import (
	"cc-backend-root-console/app/controllers"
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func (c *Controller) getHandlers() []controllers.Handler {
	return []controllers.Handler{
		c.List,
	}
}

func (c *Controller) List(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func BuildOrganizationController() *Controller {
	return &Controller{}
}
