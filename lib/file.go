package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const (
	Markdown = ".md"
	Gocode   = ".go"
)

type File struct {
	Name      string
	Path      string
	Format    string
	HasFormat bool
}

// Returns absolute file name
func (self *File) AbsoluteName() string {
	if len(self.Path) == 0 {
		self.Path, _ = self.Pwd()
	}
	return fmt.Sprintf("%v/%v", self.Path, self.Name)
}

// Gets current directory
func (self *File) Pwd() (path string, err error) {
	path, err = os.Getwd()
	return
}

// Read file and set Format field
func (f *File) Read() (data []byte, err error) {
	fmt.Printf("Path: %v \n", f.AbsoluteName())
	if data, err = ioutil.ReadFile(f.AbsoluteName()); err != nil {
		return
	}

	switch extension := path.Ext(f.Name); extension {
	case Markdown:
		f.Format = "slide"
		f.HasFormat = true
	case Gocode:
		f.Format = "code"
		f.HasFormat = true
	default:
		f.Format = ""
		f.HasFormat = false
	}
	return
}
