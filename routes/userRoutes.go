package routes

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	g := r.Group("/api/users")

	g.GET("/", controllers.ListUsers)
	g.POST("/", controllers.CreateUser)
	g.GET("/:id", controllers.GetUser)
	g.PUT("/:id", controllers.UpdateUser)
	g.DELETE("/:id", controllers.DeleteUser)

}
