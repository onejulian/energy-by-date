package dao_test

import (
	"energyByDate/dao"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDb(t *testing.T) {
	assert := assert.New(t)
	db, err := dao.ConnectPostgres()
	assert.Nil(err)
	assert.NotNil(db)
	defer db.Close()
}
