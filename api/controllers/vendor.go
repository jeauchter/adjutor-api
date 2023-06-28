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

func (server *Server) Vendors(w http.ResponseWriter, r *http.Request) {

	post := products.Vendor{}

	posts, err := post.AllVendors(server.database.Product)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}

func (server *Server) CreateVendor(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	vendor := products.Vendor{}
	err = json.Unmarshal(body, &vendor)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	vendor.PrepareVendor()
	err = vendor.ValidateVendor()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// uid, err := auth.ExtractTokenID(r)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }
	vendorCreated, err := vendor.CreateVendor(server.database.Product)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, vendorCreated.ID))
	responses.JSON(w, http.StatusCreated, vendorCreated)
}

func (server *Server) UpdateVendor(w http.ResponseWriter, r *http.Request) {

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
	vendor := products.Vendor{}
	err = server.database.Product.Debug().Model(products.Vendor{}).Where("id = ?", id).Take(&vendor).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("vendor not found"))
		return
	}

	// Read the data vendored
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	vendorUpdate := products.Vendor{}
	err = json.Unmarshal(body, &vendorUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	vendorUpdate.PrepareVendor()
	err = vendorUpdate.ValidateVendor()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	vendorUpdated, err := vendorUpdate.UpdateVendor(server.database.Product, id)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, vendorUpdated)
}

func (server *Server) GetVendorById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id64, err := strconv.ParseUint(vars["id"], 10, 64)
	id := uint32(id64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	vendor := products.Vendor{}

	vendorReceived, err := vendor.VendorById(server.database.Product, id)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, vendorReceived)
}

func (server *Server) DeleteVendor(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid vendor id given to us?
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

	// Check if the vendor exist
	vendor := products.Vendor{}
	err = server.database.Product.Debug().Model(products.Vendor{}).Where("id = ?", id).Take(&vendor).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	_, err = vendor.DeleteVendor(server.database.Product, id)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", id))
	responses.JSON(w, http.StatusNoContent, "")
}
