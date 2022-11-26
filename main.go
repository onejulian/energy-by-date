package main

import (
	"energyByDate/view"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcomen Go! Try to post a json to /generate-report ;)")
	})
	router.POST("/generate-report", view.ReportView)

	router.Run(":8186")
}
