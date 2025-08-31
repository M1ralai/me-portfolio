package server

import (
	"log"
	"net/http"

	"github.com/M1ralai/me-portfolio/cmd/api/post"
)

type Server struct {
	apiAddr string
	fsAddr  string
	apiMux  *http.ServeMux
	fsMux   *http.ServeMux
	logger  *log.Logger
}

func NewServer(apiAddr string, fsAddr string, logger *log.Logger) *Server {
	return &Server{
		apiAddr: apiAddr,
		fsAddr:  fsAddr,
		fsMux:   http.NewServeMux(),
		apiMux:  http.NewServeMux(),
		logger:  logger,
	}
}

func (s *Server) Run() {
	s.logger.Printf("server started at port %s and waiting connections", s.apiAddr)
	s.apiMux.HandleFunc("/health", s.health)
	s.apiMux.HandleFunc("/api/posts", post.HandlePostRequests)
	s.apiMux.HandleFunc("/api/posts/{id}", post.GetPostById)

	fs := http.FileServer(http.Dir(".web/"))
	s.fsMux.Handle("/", fs)

	go func() {
		s.logger.Println("html server started at port ", s.fsAddr)
		err := http.ListenAndServe(s.fsAddr, s.fsMux)
		if err != nil {
			s.logger.Fatal(err)
		}
	}()
	err := http.ListenAndServe(s.apiAddr, s.apiMux)
	if err != nil {
		s.logger.Fatal(err)
	}
}

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	s.logger.Printf("server running at %s port and health check is okay", s.apiAddr)
}
