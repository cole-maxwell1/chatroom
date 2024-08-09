package web

import (
	"net/http"

	"github.com/cole-maxwell1/chatroom/web/templates"
	"github.com/labstack/echo/v4"
)

func RenderChatRoom(c echo.Context, broker *WebSocketBroker) error {

	return Render(c, http.StatusOK, templates.ChatDisplay(broker.chatMessages.Get(), len(broker.clients)))
}
