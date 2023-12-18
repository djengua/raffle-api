package faas

import (
	"net/http"

	"github.com/djengua/rifa-api/handlers"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	handlers.TranslateHandler(w, r)
}
