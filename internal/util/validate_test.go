package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	data := struct {
		Name string `validate:"omitempty,oneof=Y N"`
	}{
		Name: "Y",
	}

	err := Validate(data)
	assert.NoError(t, err)

	tmp := struct {
		Name string `validate:"omitempty,oneof=Y N"`
	}{
		Name: "x",
	}

	err = Validate(tmp)
	assert.Error(t, err)
	t.Logf("error: %s", err)
}
