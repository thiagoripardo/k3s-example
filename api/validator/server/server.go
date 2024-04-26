package server

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thiagoripardo/k3s-example/api/validator/domain/model"
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

	s.srv.POST("/validate", func(c echo.Context) error {
		var body model.BaconRequest

		if err := c.Bind(&body); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		if !isValidBacon(body) {
			return c.JSON(http.StatusBadRequest, "invalid bacon request")
		}

		fmt.Printf("bacon pod instance: %s requested to validator instance: %s\n", body.InstanceID, instanceID.String())

		return c.String(http.StatusOK, "bacon request is valid")
	})

	s.srv.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	return s.srv.Start(":8080")
}

func isValidBacon(baconReq model.BaconRequest) bool {
	if baconReq.InstanceID == "" ||
		baconReq.Bacon.BaconName == "" ||
		baconReq.Bacon.Description == "" {
		return false
	}

	return true
}
