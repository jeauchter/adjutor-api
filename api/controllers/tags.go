package controllers

import (
	"net/http"

	"github.com/jeremyauchter/adjutor/api/responses"
	"github.com/jeremyauchter/adjutor/models/products"
)

func (server *Server) Tags(w http.ResponseWriter, r *http.Request) {

	post := products.Tag{}

	posts, err := post.AllTags(server.database.Product)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}
