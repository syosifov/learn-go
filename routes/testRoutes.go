package routes

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTestRoutes(r *gin.Engine) {

	g := r.Group("/t")

	g.GET("/hello", controllers.Hello)
	g.GET("/vars", controllers.VarsTest)

	g.POST("/forgot-password", controllers.ForgotPassword)

	g.POST("/message", controllers.TSendMessage)

}
