package env_test

import (
	"energyByDate/infraestructure/env"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	assert := assert.New(t)
	assert.NotEmpty(env.Env("DB_HOST"))
}
