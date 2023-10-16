package server

import (
	"doodocsTask/config"
	"doodocsTask/pkg/logger"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(config config.Config, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           config.Host + ":" + config.Port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	logger.InfoLog.Printf("Server run on http://localhost:%s", config.Port)
	return s.httpServer.ListenAndServe()
}
