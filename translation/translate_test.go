package translation_test

import (
	"testing"

	"github.com/djengua/raffle-api/translation"
)

func TestTranslate(t *testing.T) {
	// Arrange | Disponer
	tt := []struct {
		Word        string
		Language    string
		Translation string
	}{
		{
			Word:        "hello",
			Language:    "english",
			Translation: "hello",
		},
		{
			Word:        "hello",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello",
			Language:    "finnish",
			Translation: "hei",
		},
		{
			Word:        "hello",
			Language:    "dutch",
			Translation: "",
		},
		{
			Word:        "hello",
			Language:    "spanish",
			Translation: "hola",
		},
		{
			Word:        "bye",
			Language:    "spanish",
			Translation: "hola",
		},
	}

	for _, test := range tt {
		// Act | Actuar
		underTest := translation.NewStaticService()
		res := underTest.Translate(test.Word, test.Language)

		// Assert | Afirmar
		if res != test.Translation {
			t.Errorf(`expected "%s" to be "%s" from "%s" but received "%s"`, test.Word, test.Language, test.Translation, res)
		}
	}
}
