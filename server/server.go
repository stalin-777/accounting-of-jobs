package server

import (
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stalin-777/accounting-of-jobs/config"
	"github.com/stalin-777/accounting-of-jobs/http"
	"github.com/stalin-777/accounting-of-jobs/logger"
	"github.com/stalin-777/accounting-of-jobs/postgres"
)

type Server struct {
	Router *echo.Echo

	WorkplaceService postgres.WorkplaceService
}

func Run(cfg *config.Config, connPool *pgxpool.Pool) {

	server, err := getRouter(cfg, connPool)
	if err != nil {
		logger.Fatalf("Failed to ger router, error:%s", err.Error())
	}

	routerSocket := fmt.Sprintf("%s:%v", cfg.Web.Host, cfg.Web.Port)
	logger.Infof("Service is running on socket %v\n", routerSocket)
	if err := server.Router.Start(routerSocket); err != nil {
		logger.Fatalf("Failed to start service, error:%s", err.Error())
	}
}

func getRouter(cfg *config.Config, connPool *pgxpool.Pool) (*Server, error) {

	s := &Server{
		Router:           echo.New(),
		WorkplaceService: postgres.WorkplaceService{DB: connPool},
	}
	s.registerHandlers()

	s.Router.Use(middleware.Recover())

	return s, nil
}

func (s *Server) registerHandlers() {

	var h http.Handler

	s.Router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "API status: online")
	})

	h.WorkplaceService = &s.WorkplaceService

	s.Router.GET("/workplaces/:id", h.FindWorkplace)
	s.Router.GET("/workplaces", h.FindWorkplaces)
	s.Router.POST("/workplaces", h.CreateWorkplace)
	s.Router.PUT("/workplaces/:id", h.UpdateWorkplace)
	s.Router.DELETE("/workplaces/:id", h.DeleteWorkplace)
}
