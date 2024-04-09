package server

import (
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

type Server struct {
	srv *echo.Echo
}

var (
	instance *Server
	once     sync.Once
)

func New() *Server {
	once.Do(func() {
		instance = &Server{
			srv: echo.New(),
		}
	})

	return instance
}

func (s *Server) Run() error {
	s.srv.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "echo v4 server")
	})

	s.srv.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	return s.srv.Start(":8080")
}
