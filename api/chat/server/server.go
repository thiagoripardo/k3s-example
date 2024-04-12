package server

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thiagoripardo/k3s-example/api/chat/domain/model"
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
	instanceID := uuid.New()

	s.srv.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to lovelybacon.com!")
	})

	s.srv.GET("/generator", func(c echo.Context) error {
		baconRes := model.BaconResponse{
			InstanceID: instanceID.String(),
		}

		randomInt := rand.Intn(50)
		bacon := model.Bacons[randomInt%len(model.Bacons)]
		baconRes.Bacon = bacon

		fmt.Printf("instance: %s generated a bacon\n", baconRes.InstanceID)

		return c.JSON(http.StatusOK, baconRes)
	})

	s.srv.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	return s.srv.Start(":8080")
}
