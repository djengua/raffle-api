package translation

import "strings"

func Translate(word string, language string) string {
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
