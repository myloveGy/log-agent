package tail

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"log-agent/util"
)

func TestTailHandle_toFilename(t *testing.T) {
	tailHandle := &TailHandle{config: &Config{Filename: "laravel-2006-01-02.log", Rule: true}}
	assert.Equal(t, fmt.Sprintf("/laravel-%s.log", util.Date()), tailHandle.toFilename())

	tailHandle = &TailHandle{config: &Config{Filename: "wechat-2006-01-02.log", Rule: true}}
	assert.Equal(t, fmt.Sprintf("/wechat-%s.log", util.Date()), tailHandle.toFilename())
}
