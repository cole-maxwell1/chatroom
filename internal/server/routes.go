package server

import (
	"log"
	"net/http"
	"time"

	"github.com/cole-maxwell1/chatroom/web"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      skipWebSocket,
		ErrorMessage: "Request timeout",
		OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
			log.Println(c.Path())
		},
		Timeout: 10 * time.Second,
	}))

	// Serve static files
	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/javascript/*", echo.WrapHandler(fileServer))
	e.GET("/css/main.css", echo.WrapHandler(fileServer))

	broker := web.NewHub()
	go broker.Run()
	e.GET("/ws", func(c echo.Context) error {
		web.HandleWebSocket(broker, c.Response(), c.Request())
		return nil
	})

	// Register routes
	e.GET("/", func(c echo.Context) error {
		err := web.RenderChatRoom(c, broker)
		if err != nil {
			return err
		}
		return nil
	})

	return e
}

func skipWebSocket(c echo.Context) bool {
	// Skip middleware if path is equal 'login'
	if c.Request().URL.Path == "/ws" {
		return true
	}
	return false
}
