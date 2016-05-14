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
	s.Render = render.New(render.Options{
		Layout:        "layout",
		IsDevelopment: true,
	})
	http.ListenAndServe(":"+s.Port, http.HandlerFunc(s.ServeHTTP))
	return
}

func (s *Server) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Printf("URL: %v", r.URL)
	var file = File{Name: r.URL.Path[1:]}
	log.Printf("URL: %v", r.URL.Path)
	if r.URL.Path == "/" {
		file.Name = "README.md"
	}
	var result interface{}
	var err error
	if result, err = file.Read(); err != nil {
		http.Error(rw, "Unable to read file", 500)
		return
	}
	if file.HasFormat {
		s.Render.HTML(rw, 200, file.Format, result)
		return
	}

	fs := http.FileServer(http.Dir("."))
	fs.ServeHTTP(rw, r)
}
