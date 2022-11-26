package formatDate_test

import (
	"energyByDate/infraestructure/formatDate"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatToEn(t *testing.T) {
	assert := assert.New(t)
	date, err := formatDate.FormatToEn("octubre 25, 2022, 12:50 AM")
	assert.Nil(err)
	assert.NotEmpty(date)
}
