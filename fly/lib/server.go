package lib

import (
	"github.com/theplant/blackfriday"
	"gopkg.in/unrolled/render.v1"
	"html/template"
	"log"
	"net/http"
)

type Server struct {
	Debug bool
	Index string
	Port  string
	Path  string
	*render.Render
}

func (s *Server) Run() (err error) {
	log.Println("Listening on port", s.Port)
	s.Render = render.New(render.Options{
		Layout:        "layout",
		IsDevelopment: true,
	})
	http.ListenAndServe(":"+s.Port, http.HandlerFunc(s.ServeHTTP))
	return
}

func (s *Server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if s.Debug {
		log.Printf("URL: %v", r.URL)
	}
	var file = File{Name: r.URL.Path[1:]}

	if r.URL.Path == "/" {
		file.Name = s.Index
	}
	file.Path = s.Path
	var data []byte
	var err error
	if data, err = file.Read(); err != nil {
		http.Error(rw, "Unable to read file", 500)
		return
	}
	if file.HasFormat {
		s.ShowMarkdown(rw, data, file)
		return
	}

	s.Show(rw, r)
}

func (s *Server) ShowMarkdown(rw http.ResponseWriter, data []byte, file File) {
	md := blackfriday.MarkdownCommon(data)
	var result = template.HTML(md)
	s.Render.HTML(rw, 200, file.Format, result)
}

func (s *Server) Show(rw http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("."))
	fs.ServeHTTP(rw, r)
}
