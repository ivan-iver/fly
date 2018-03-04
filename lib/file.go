package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const (
	// Markdown file extension
	Markdown = ".md"
	// Gocode is a file extension constant
	Gocode = ".go"
)

// File models a file into file system
type File struct {
	Name      string
	Path      string
	Format    string
	HasFormat bool
}

// AbsoluteName returns absolute file name
func (f *File) AbsoluteName() string {
	if len(f.Path) == 0 {
		f.Path, _ = f.Pwd()
	}
	return fmt.Sprintf("%v/%v", f.Path, f.Name)
}

// Pwd gets current directory
func (f *File) Pwd() (path string, err error) {
	path, err = os.Getwd()
	return
}

// Exists identify if the file exist
func (f *File) Exists() (exists bool, err error) {
	if _, err = os.Stat(f.AbsoluteName()); err == nil {
		exists = true
	}
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
