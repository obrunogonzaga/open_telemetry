package webserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

func (s *WebServer) Start() error {

	// Middlewares
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.Timeout(60 * time.Second))

	for path, handler := range s.Handlers {
		s.Router.Handle(path, handler)
	}

	// Register Prometheus metrics handler
	s.Router.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(s.WebServerPort, s.Router)
	if err != nil {
		return err
	}
	return nil
}
