package server

import (
	"net/http"
	"log"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func Init() {
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

	r.POST("/signup", func(c *gin.Context) {
		log.Println("exec signup function")
		c.JSON(http.StatusOK, "foo")
	})
	r.Run(":8080")
}
