package lib

import (
	"io/ioutil"
	"path"
)

const (
	Markdown = ".md"
	Gocode   = ".go"
)

type File struct {
	Name      string
	Format    string
	HasFormat bool
}

func (f *File) Read() (data []byte, err error) {
	if data, err = ioutil.ReadFile(f.Name); err != nil {
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
