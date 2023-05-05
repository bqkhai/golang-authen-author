package routes

import (
	"authen-author-example/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("api/auth/signup", controllers.Signup)
	r.POST("api/auth/login", controllers.Login)
}
