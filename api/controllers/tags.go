package controllers

import (
	"net/http"

	"github.com/jeremyauchter/adjutor/api/responses"
)

func (server *Server) Tags(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "This is Tags")

}
