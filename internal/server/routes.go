package server

import (
	"net/http"

	"github.com/cole-maxwell1/chatroom/web"
	"github.com/cole-maxwell1/chatroom/web/templates"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/javascript/*", echo.WrapHandler(fileServer))
	e.GET("/css/main.css", echo.WrapHandler(fileServer))

	e.GET("/", echo.WrapHandler(templ.Handler(templates.HelloForm())))
	e.POST("/hello", echo.WrapHandler(http.HandlerFunc(web.HelloWebHandler)))

	return e
}



