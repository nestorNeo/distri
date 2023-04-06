package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nestorneo/distri/middleware"
)

func main() {

	r := gin.Default()
	r.Use(middleware.GuidMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run("0.0.0.0:5000")
}

/*

	Microservicios
	  3 - (cluster) [servicio A] - api (server) + client
	  [servicio B]
	  [servicio C]


*/
