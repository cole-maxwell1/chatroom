package web

import (
	"net/http"

	"github.com/cole-maxwell1/chatroom/web/templates"
)

func HelloWebHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	name := r.FormValue("name")
	component := templates.HelloPost(name)
	component.Render(r.Context(), w)
}
