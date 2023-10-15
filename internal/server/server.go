package server

import (
	"doodocsTask/pkg/logger"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler gin.IRoutes) error {
	s.httpServer = &http.Server{
		Addr:           port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	logger.InfoLog.Printf("Server run on http://localhost:%s", port)
	return s.httpServer.ListenAndServe()
}
