package lib

import (
	"gopkg.in/unrolled/render.v1"
	"log"
	"net/http"
)

type Server struct {
	Port string
	*render.Render
}

func (s *Server) Run() (err error) {
	log.Println("Listening on port", s.Port)
	http.ListenAndServe(":"+s.Port, http.HandlerFunc(s.ServeHTTP))
	return
}

func (s *Server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var file = File{Name: r.URL.Path[1:]}
	if r.URL.Path == "/" {
		file.Name = "README.md"
	}
	var result string
	var err error
	if result, err = file.Read(); err != nil {
		http.Error(rw, "Unable to read file", 500)
		return
	}
	if file.IsMarkdown {
		s.Render.HTML(rw, 200, "slide", result)
		return
	}

	fs := http.FileServer(http.Dir("."))
	fs.ServeHTTP(rw, r)
}
