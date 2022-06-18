package util

import (
	"path/filepath"
	"runtime"
)

func TestDataPath(path ...string) string {
	_, file, _, _ := runtime.Caller(0)
	paths := []string{filepath.Dir(filepath.Dir(file)), "testdata"}
	paths = append(paths, path...)
	return filepath.Join(paths...)
}
