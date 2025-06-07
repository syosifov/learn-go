package routes

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(r *gin.Engine) {

	g := r.Group("/admin")
	g.GET("/hello", controllers.Hello)

}
