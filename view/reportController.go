package view

import (
	usecase "energyByDate/useCase"
	formrequest "energyByDate/view/formRequest"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateReport(c *gin.Context) {
	var req formrequest.FormRequest
	var getReportUseCase usecase.GetReportUseCase
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	report, err := getReportUseCase.Execute(req.Date, req.Period)
	if err != nil {
		if err.Error() == "period not found" {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSONP(http.StatusOK, report)
}
