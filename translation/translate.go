package translation

import "strings"

// type Translator interface {
// 	Translate(word string, language string) string
// }

// type TranslateHandler struct {
// 	service Translator
// }

// func NewTranslateHandler(service Translator) *TranslateHandler {
// 	return &TranslateHandler{
// 		service: service,
// 	}
// }

type StaticService struct{}

func NewStaticService() *StaticService {
	return &StaticService{}
}

func (s *StaticService) Translate(word string, language string) string {
	lng := sanitizeInput(language)
	switch lng {
	case "english":
		return "hello"
	case "finnish":
		return "hei"
	case "german":
		return "hallo"
	case "spanish":
		return "hola"
	default:
		return ""
	}
}

func sanitizeInput(w string) string {
	w = strings.ToLower(w)
	return strings.TrimSpace(w)
}
