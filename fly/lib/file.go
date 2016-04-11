package lib

import (
	"github.com/theplant/blackfriday"
	"html/template"
	"io/ioutil"
	"strings"
)

type File struct {
	Name       string
	IsMarkdown bool
}

func (f *File) Read() (result string, err error) {
	var data []byte
	if data, err = ioutil.ReadFile(f.Name); err != nil {
		return
	}
	if strings.HasSuffix(f.Name, ".md") {
		f.IsMarkdown = true
		var md = blackfriday.MarkdownCommon(data)
		result = string(template.HTML(md))
	} else {
		result = string(data)
	}
	return
}
