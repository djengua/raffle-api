package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/djengua/rifa-api/translation"
)

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

const defaultLanguage = "english"

func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	language := r.URL.Query().Get("language")
	if language == "" {
		language = defaultLanguage
	}
	word := strings.ReplaceAll(r.URL.Path, "/", "")
	translation := translation.Translate(word, language)
	response := Resp{
		Language:    language,
		Translation: translation,
	}
	if err := enc.Encode(response); err != nil {
		panic("unable to encode response")
	}
}
