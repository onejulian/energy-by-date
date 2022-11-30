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
		c.String(http.StatusOK, "Welcome! Try to open a bash terminal and run the following command: curl -X POST https://energy-by-date.salmonisland-c75248cf.eastus.azurecontainerapps.io/generate-report -H 'Content-Type: application/json' -d '{\"date\":\"2022-10-25\",\"period\":\"daily\"}'")
	})
	router.POST("/generate-report", view.ReportView)

	router.Run(":8186")
}
