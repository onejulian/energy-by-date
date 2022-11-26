package translatemonth_test

import (
	translatemonth "energyByDate/infraestructure/translateMonth"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslateEnToSp(t *testing.T) {
	assert := assert.New(t)
	month, err := translatemonth.TranslateEnToSp("January")
	assert.Nil(err)
	assert.Equal("enero", month)
}

func TestTranslateSpToEn(t *testing.T) {
	assert := assert.New(t)
	month, err := translatemonth.TranslateSpToEn("enero")
	assert.Nil(err)
	assert.Equal("January", month)
}
