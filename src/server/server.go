package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (srv *Server) Run() error {
	r := mux.NewRouter()

	createRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s\n", port)
	}

	s := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	log.Printf("Listening on port %s\n", port)
	return s.ListenAndServe()
}

func createRoutes(r *mux.Router) {
	// // Games
	// gRouter := r.PathPrefix("/games").Subrouter()

	// gRouter.PathPrefix("/").HandlerFunc(notFound)

	// PaperCut bug endpoint
	r.Methods("GET").Path("/pc-bug").HandlerFunc(mockDownloadInfo)
	r.Methods("POST").Path("/pc-bug").HandlerFunc(logResponse)

	// 404 Handler
	r.PathPrefix("/").HandlerFunc(notFound)
}
