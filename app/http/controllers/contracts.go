package controllers

import "github.com/gin-gonic/gin"

type Handler func(ctx *gin.Context)

type Controller interface {
	getHandlers() []*Handler
}
