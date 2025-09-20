package routes

import (
	_ "ecosnap/docs"
	"ecosnap/internal/controllers"
	"ecosnap/internal/controllers/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Static("/uploads", "./uploads")

	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", controllers.SignUp)
		auth.POST("/sign-in", controllers.SignIn)
		auth.POST("/refresh", controllers.RefreshToken)
	}

	r.GET("/users", controllers.GetAllUsers)
	r.GET("/users/:id", controllers.GetUserByID)

	profile := r.Group("/profile", middlewares.CheckUserAuthentication)
	{
		profile.GET("", controllers.GetMyDataUser)
		profile.PATCH("", controllers.UpdateUser)
		profile.PATCH("/password", controllers.UpdateUsersPassword)
	}

	return r
}
