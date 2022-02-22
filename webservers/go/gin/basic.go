package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gsainz/autocannon/webservers/go"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, weather.Predict(5))
	})

	return r
}

func main() {
	r := setupRouter()
	log.Fatal(r.Run(":3000"))
}
