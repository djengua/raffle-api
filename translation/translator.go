package translation

func Translate(word string, language string) string {
	switch language {
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
