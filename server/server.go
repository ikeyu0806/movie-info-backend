package server

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/ikeyu0806/movie-info-backend/controllers"
)

func Init() {
	r := router()

	r.Run(":8080")
}

func router() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{

		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},

		AllowHeaders: []string{
			"Access-Control-Allow-Origin",			
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},

		AllowOrigins: []string{
			"http://localhost:3001",
		},
		MaxAge: 24 * time.Hour,
	}))

	user_ctrl := user.Controller{}
	r.POST("/signup", user_ctrl.SignUp)
	r.POST("/login", user_ctrl.Login)

	return r
}
