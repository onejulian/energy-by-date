package view

import (
	usecase "energyByDate/useCase"
	formrequest "energyByDate/view/formRequest"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReportView(c *gin.Context) {
	var req formrequest.FormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	report, err := usecase.GetReportUseCase(req.Date, req.Period)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, report)
}
