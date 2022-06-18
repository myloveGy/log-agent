package util

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(map[string]string{"username": "jinxing.liu"}, "secret123456", 24*time.Hour)
	assert.NoError(t, err)
	fmt.Println(token)

	item := make(map[string]string)
	_, err = ParseToken(&item, token, "secret123456")
	assert.NoError(t, err)
	assert.Equal(t, "jinxing.liu", item["username"])

	_, err = ParseToken(nil, token, "secret1234568999")
	assert.Error(t, err)
	t.Logf("error: %v", err)
}
