package lib_test

import (
	"fmt"
	"github.com/iver/fly/lib"
	"testing"
)

func TestAbsoluteName(t *testing.T) {
	path := "/Users/iver/Workspace/go/src/github.com/iver/fly/lib"
	file := "file.go"

	t.Run("AbsoluteName()", func(t *testing.T) {
		var f = lib.File{
			Name: file,
			Path: path,
		}
		fullname := fmt.Sprintf("%v/%v", path, file)
		//	t.Logf("Fullname: %v", fullname)
		//	t.Logf("Absolute: %v", f.AbsoluteName())
		if f.AbsoluteName() != fullname {
			t.Error("AbsoluteName is different")
		}
	})

}
