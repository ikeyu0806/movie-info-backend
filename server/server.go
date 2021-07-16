package server

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ikeyu0806/movie-info-backend/controllers"
)

func Init() {
	r := router()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	port = ":" + port
	r.Run(port)
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

	user_ctrl := controllers.UserController{}
	review_ctrl := controllers.ReviewController{}
	running_ctrl := controllers.RunningController{}
	r.GET("/is_running", running_ctrl.IsRunning)
	r.POST("/signup", user_ctrl.SignUp)
	r.POST("/login", user_ctrl.Login)
	r.GET("/user/:name", user_ctrl.FindByName)
	r.POST("/review/create", review_ctrl.Create)
	r.GET("/review/:movie_id", review_ctrl.FindByMovieID)

	return r
}
