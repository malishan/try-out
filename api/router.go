package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/try-out/api/v1"
)

func GetRouter(middleware ...gin.HandlerFunc) *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())

	group := router.Group("notification")

	v1Group := group.Group("/v1")
	v1.NotificationHandler(v1Group)

	return router
}
