package util

import (
	"fmt"
	"testing"
)

func TestValidate(t *testing.T) {
	data := struct {
		Name string `validate:"omitempty,oneof=Y N"`
	}{
		Name: "Y",
	}

	err := Validate(data)
	fmt.Println("error", err)
}
