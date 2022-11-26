package domain_test

import (
	"energyByDate/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReport(t *testing.T) {
	assert := assert.New(t)
	report := domain.Report{}
	report, err := report.GetReport("2022-10-25", "monthly")
	assert.Nil(err)
	assert.NotNil(report)
}
