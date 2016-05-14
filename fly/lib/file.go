package lib

import (
	"github.com/theplant/blackfriday"
	"html/template"
	"io/ioutil"
	"log"
	"strings"
)

type File struct {
	Name      string
	Format    string
	HasFormat bool
}

func (f *File) Read() (result interface{}, err error) {
	var data []byte
	if data, err = ioutil.ReadFile(f.Name); err != nil {
		return
	}
	if strings.HasSuffix(f.Name, ".md") {
		f.Format = "slide"
		f.HasFormat = true
		log.Printf("data.md: slide")
		var md = blackfriday.MarkdownCommon(data)
		result = template.HTML(md)
	} else if strings.HasSuffix(f.Name, ".go") {
		f.Format = "code"
		f.HasFormat = true
		log.Printf("data.md: code")
		var md = blackfriday.MarkdownCommon(data)
		result = template.HTML(md)
	} else {
		f.Format = ""
		log.Printf("data.algo mas: %s", data)
		result = string(data)
	}
	return
}
