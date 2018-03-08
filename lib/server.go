package lib

import (
	"github.com/russross/blackfriday"
	"gopkg.in/unrolled/render.v1"
	"html/template"
	"net/http"
)

// Server represents web server logic and attributes
type Server struct {
	Debug bool
	Index string
	Port  string
	Path  string
	*render.Render
}

// Run publish the web server
func (s *Server) Run() (err error) {
	log.Info("Listening on port", s.Port)
	s.Render = render.New(render.Options{
		Layout:        "layout",
		IsDevelopment: true,
	})
	http.ListenAndServe(":"+s.Port, http.HandlerFunc(s.ServeHTTP))
	return
}

// ServerHTTP is the handler function to publish into web server
func (s *Server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if s.Debug {
		log.Debug("URL:", r.URL)
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

// ShowMarkdown contains generation logic from markdown to html.
func (s *Server) ShowMarkdown(rw http.ResponseWriter, data []byte, file File) {
	md := blackfriday.MarkdownCommon(data)
	var result = template.HTML(md)
	s.Render.HTML(rw, 200, file.Format, result)
}

// Show is the function that publish all files in the directory.
func (s *Server) Show(rw http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("."))
	fs.ServeHTTP(rw, r)
}
