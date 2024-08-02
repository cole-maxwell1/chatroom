package web

/*
import (
	"bytes"
	"context"
	"log"

	"github.com/cole-maxwell1/chatroom/web/templates"
	"github.com/labstack/echo/v4"
	"github.com/r3labs/sse/v2"
)

func HandleServerSentEvents(c echo.Context) error { // longer variant with disconnect logic
	server := sse.New()       // create SSE broadcaster server
	server.AutoReplay = false // do not replay messages for each new subscriber that connects
	server.CreateStream("new-message")

	go func(s *sse.Server, ctx context.Context) {
		// listen to chatChan for new messages
		for msg := range ChatChan {
			// send the new message to the client
			buf := new(bytes.Buffer)

			// Render chat template
			err := templates.ChatMessage(msg).Render(ctx, buf)
			if err != nil {
				log.Printf("error rendering chat message: %v\n", err)
			}

			// Send the message to the client
			s.Publish("new-message", &sse.Event{
				Data: buf.Bytes(),
			})
		}
	}(server, c.Request().Context())

	server.ServeHTTP(c.Response(), c.Request())
	return nil

}
 */