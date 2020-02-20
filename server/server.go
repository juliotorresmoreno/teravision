package server

import (
	"net/http"

	"github.com/juliotorresmoreno/teravision/config"
	"github.com/juliotorresmoreno/teravision/controllers"
	"github.com/labstack/echo"
)

// Server http service
type Server struct {
	*echo.Echo
}

// Listen raise a new server
func (e Server) Listen() {
	conf := config.GetConf()
	host := conf.GetFullHost()
	e.Logger.Fatal(e.Start(host))
}

// NewServer configure server, middlewares and routes
func NewServer() *Server {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	api := e.Group("/api")
	controllers.RouteUsers(api.Group("/users"))

	return &Server{e}
}
