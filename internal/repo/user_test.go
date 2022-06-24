package repo

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestUserRepo_Pagination(t *testing.T) {
	items := make([]map[string]interface{}, 0)
	value := reflect.ValueOf(&items)

	// json.Unmarshal returns errors for these
	if value.Kind() != reflect.Ptr {
		t.Log("must pass a pointer, not a value, to StructScan destination")
	}

	if value.IsNil() {
		t.Log("nil pointer passed to StructScan destination")
	}

	valueType := value.Type()
	if valueType.Kind() == reflect.Ptr {
		valueType = valueType.Elem()
	}

	if valueType.Kind() != reflect.Slice {
		t.Log("must pass a pointer to a slice, not a value, to StructScan destination")
	}

	isPtr := valueType.Elem().Kind() == reflect.Ptr
	base := valueType.Elem()
	if base.Kind() == reflect.Ptr {
		base = base.Elem()
	}

	direct := reflect.Indirect(value)

	vp := reflect.New(base)
	json.Unmarshal([]byte(`{"username":"jinxing.liu"}`), vp.Interface())

	// append
	if isPtr {
		direct.Set(reflect.Append(direct, vp))
	} else {
		direct.Set(reflect.Append(direct, reflect.Indirect(vp)))
	}

	fmt.Println(direct, items)
}
