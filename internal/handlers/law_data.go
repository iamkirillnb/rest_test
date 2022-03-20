package handlers

import (
	"fmt"
	Openapi "github.com/iamkirillnb/rus_law/internal/oapi/law"
	"github.com/iamkirillnb/rus_law/pkg/logging"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const (
	defaultServerReadTimeout  = 15 * time.Second
	defaultServerWriteTimeout = 30 * time.Second
)

type server struct {
	*http.Server

	log logging.Logger
}

func NewHandler(logger logging.Logger) *server {
	ser := &http.Server{
		Addr:         ":8080",
		Handler:      echo.New(),
		ReadTimeout:  defaultServerReadTimeout,
		WriteTimeout: defaultServerWriteTimeout,
	}
	a := &server{
		Server: ser,
		log:    logger,
	}

	Openapi.RegisterHandlers(echo.New(), a)

	return a
}

func (s *server) GetFormData(ctx echo.Context) error {

	return nil
}

func (s *server) PostFormData(ctx echo.Context) error {
	fmt.Println("Hello POST method")
	return nil
}
