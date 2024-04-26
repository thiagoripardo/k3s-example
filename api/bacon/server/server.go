package server

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thiagoripardo/k3s-example/api/bacon/domain/model"
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

		isValid, err := verifyGeneratedBacon(baconRes)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		if !isValid {
			return c.JSON(http.StatusBadRequest, "invalid bacon request")
		}

		fmt.Printf("instance: %s generated a bacon\n", baconRes.InstanceID)

		return c.JSON(http.StatusOK, baconRes)
	})

	s.srv.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	return s.srv.Start(":8080")
}

func verifyGeneratedBacon(baconRes model.BaconResponse) (bool, error) {
	payload, err := baconRes.AsBuffer()
	if err != nil {
		return false, err
	}

	baseURL := "http://validator.default.svc.cluster.local:3001"
	url := fmt.Sprintf("%s/validate", baseURL)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return false, fmt.Errorf("invalid bacon request")
	}

	return true, nil
}
