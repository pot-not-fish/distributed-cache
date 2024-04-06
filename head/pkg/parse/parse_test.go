package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	Init("../../config.yaml")

	assert.EqualValues(t, ConfigStructure.Admin.Username, "likea")
	assert.EqualValues(t, ConfigStructure.Admin.Password, "123456")
}
