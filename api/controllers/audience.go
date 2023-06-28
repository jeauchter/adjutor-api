package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jeremyauchter/adjutor/api/responses"
	"github.com/jeremyauchter/adjutor/models/products"
	"github.com/jeremyauchter/adjutor/utils/formaterror"
)

func (server *Server) Audiences(w http.ResponseWriter, r *http.Request) {

	post := products.Audience{}

	posts, err := post.AllAudiences(server.database.Product)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}

func (server *Server) CreateAudience(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	audience := products.Audience{}
	err = json.Unmarshal(body, &audience)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	audience.PrepareAudience()
	err = audience.ValidateAudience()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// uid, err := auth.ExtractTokenID(r)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }
	audienceCreated, err := audience.CreateAudience(server.database.Product)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, audienceCreated.ID))
	responses.JSON(w, http.StatusCreated, audienceCreated)
}

func (server *Server) UpdateAudience(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the post id is valid
	id64, err := strconv.ParseUint(vars["id"], 10, 32)
	id := uint32(id64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//CHeck if the auth token is valid and  get the user id from it
	// uid, err := auth.ExtractTokenID(r)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }

	// Check if the post exist
	audience := products.Audience{}
	err = server.database.Product.Debug().Model(products.Audience{}).Where("id = ?", id).Take(&audience).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("audience not found"))
		return
	}

	// Read the data audienceed
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	audienceUpdate := products.Audience{}
	err = json.Unmarshal(body, &audienceUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	audienceUpdate.PrepareAudience()
	err = audienceUpdate.ValidateAudience()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	audienceUpdated, err := audienceUpdate.UpdateAudience(server.database.Product, id)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, audienceUpdated)
}

func (server *Server) GetAudienceById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id64, err := strconv.ParseUint(vars["id"], 10, 64)
	id := uint32(id64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	audience := products.Audience{}

	audienceReceived, err := audience.AudienceById(server.database.Product, id)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, audienceReceived)
}

func (server *Server) DeleteAudience(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid audience id given to us?
	id64, err := strconv.ParseUint(vars["id"], 10, 64)
	id := uint32(id64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// // Is this user authenticated?
	// uid, err := auth.ExtractTokenID(r)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }

	// Check if the audience exist
	audience := products.Audience{}
	err = server.database.Product.Debug().Model(products.Audience{}).Where("id = ?", id).Take(&audience).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	_, err = audience.DeleteAudience(server.database.Product, id)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", id))
	responses.JSON(w, http.StatusNoContent, "")
}
