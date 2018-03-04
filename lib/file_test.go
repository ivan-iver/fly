package lib_test

import (
	"fmt"
	"github.com/iver/fly/lib"
	"testing"
)

func TestFile(t *testing.T) {
	path := "/Users/iver/Workspace/go/src/github.com/iver/fly/lib"
	file := "file.go"
	var f = lib.File{
		Name: file,
		Path: path,
	}

	t.Run("AbsoluteName()", func(t *testing.T) {
		fullname := fmt.Sprintf("%v/%v", path, file)
		//	t.Logf("Fullname: %v", fullname)
		//	t.Logf("Absolute: %v", f.AbsoluteName())
		if f.AbsoluteName() != fullname {
			t.Error("AbsoluteName is different")
		}
	})

	t.Run("Exists()", func(t *testing.T) {

	})

}
