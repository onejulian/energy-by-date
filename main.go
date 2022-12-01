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
		c.String(http.StatusOK, "Welcome! Try to open a bash terminal and run the following command: curl -X POST https://energy-z52mtgqafq-uc.a.run.app/generate-report -H 'Content-Type: application/json' -d '{\"date\":\"2022-10-25\",\"period\":\"daily\"}'")
	})
	router.POST("/generate-report", view.GenerateReport)

	router.Run(":8186")
}
