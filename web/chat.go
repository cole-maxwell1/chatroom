package web

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/cole-maxwell1/chatroom/web/templates"
	"github.com/labstack/echo/v4"
)

func RenderChatRoom(c echo.Context) error {
	return Render(c, http.StatusOK, templates.ChatDisplay(chatMessages.Get()))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
