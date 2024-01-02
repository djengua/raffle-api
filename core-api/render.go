package coreapi

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./templates"),
	jet.InDevelopmentMode(),
)

func HomePage(w http.ResponseWriter, r http.Request) {
	err := renderPage(w, "home", nil)
	if err != nil {
		log.Println(err)
	}
}

// func (s *Server) homePage(ctx *gin.Context) {
// 	raffle, err := s.fetchAll(ctx)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, map[string]interface{}{
// 		"error": nil,
// 		"data":  raffle,
// 	})
// }

func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println(err)
		return err
	}
	err = view.Execute(w, data, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
