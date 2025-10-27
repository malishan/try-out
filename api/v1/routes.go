package v1

import "github.com/gin-gonic/gin"

func NotificationHandler(r *gin.RouterGroup) {
	r.POST("/create", createNotification)
	r.POST("/mail", sendNotification)
	r.GET("/list", listNotification)
	r.DELETE("/", deleteNotification)
}
