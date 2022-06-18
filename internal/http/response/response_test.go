package response

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInvalidParameter(t *testing.T) {
	err := NewInvalidParameter("invalid parameter")
	assert.Equal(t, "invalid parameter", err.Error())

	err = NewInvalidParameter(fmt.Errorf("invalid parameter"))
	assert.Equal(t, "invalid parameter", err.Error())
}

func TestNewUnauthenticated(t *testing.T) {
	err := NewUnauthenticated("unauthenticated")
	assert.Equal(t, "unauthenticated", err.Error())

	err = NewUnauthenticated(fmt.Errorf("unauthenticated"))
	assert.Equal(t, "unauthenticated", err.Error())
}

func TestNewSystemError(t *testing.T) {
	err := NewSystemError("system error")
	assert.Equal(t, "system error", err.Error())

	err = NewSystemError(fmt.Errorf("system error"))
	assert.Equal(t, "system error", err.Error())
}
