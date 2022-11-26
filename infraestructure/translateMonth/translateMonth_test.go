package translatemonth_test

import (
	translatemonth "energyByDate/infraestructure/translateMonth"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslate(t *testing.T) {
	assert := assert.New(t)
	month, err := translatemonth.Translate("January")
	assert.Nil(err)
	assert.Equal("enero", month)
}
