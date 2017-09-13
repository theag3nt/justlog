package api

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// Server api server
type Server struct {
	port    string
	logPath string
}

// NewServer create Server
func NewServer(port string, logPath string) Server {
	return Server{
		port:    port,
		logPath: logPath,
	}
}

// Init api server
func (s *Server) Init() {

	e := echo.New()
	//e.HideBanner = true

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/channel/:channel/user/:username", s.getCurrentUserLogs)
	e.GET("/channel/:channel", s.getCurrentChannelLogs)
	e.GET("/channel/:channel/:year/:month/:day", s.getDatedChannelLogs)
	e.GET("/channel/:channel/user/:username/:year/:month", s.getDatedUserLogs)
	e.GET("/channel/:channel/user/:username/random", s.getRandomQuote)

	log.Printf("starting API on port :%s", s.port)
	e.Logger.Fatal(e.Start(":" + s.port))
}
