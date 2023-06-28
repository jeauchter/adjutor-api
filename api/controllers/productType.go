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

func (server *Server) ProductTypes(w http.ResponseWriter, r *http.Request) {

	post := products.ProductType{}

	posts, err := post.AllProductTypes(server.database.Product)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}

func (server *Server) CreateProductType(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	productType := products.ProductType{}
	err = json.Unmarshal(body, &productType)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	productType.PrepareProductType()
	err = productType.ValidateProductType()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// uid, err := auth.ExtractTokenID(r)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }
	productTypeCreated, err := productType.CreateProductType(server.database.Product)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, productTypeCreated.ID))
	responses.JSON(w, http.StatusCreated, productTypeCreated)
}

func (server *Server) UpdateProductType(w http.ResponseWriter, r *http.Request) {

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
	productType := products.ProductType{}
	err = server.database.Product.Debug().Model(products.ProductType{}).Where("id = ?", id).Take(&productType).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("productType not found"))
		return
	}

	// Read the data productTypeed
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	productTypeUpdate := products.ProductType{}
	err = json.Unmarshal(body, &productTypeUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	productTypeUpdate.PrepareProductType()
	err = productTypeUpdate.ValidateProductType()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	productTypeUpdated, err := productTypeUpdate.UpdateProductType(server.database.Product, id)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, productTypeUpdated)
}

func (server *Server) GetProductTypeById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id64, err := strconv.ParseUint(vars["id"], 10, 64)
	id := uint32(id64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	productType := products.ProductType{}

	productTypeReceived, err := productType.ProductTypeById(server.database.Product, id)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, productTypeReceived)
}

func (server *Server) DeleteProductType(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid productType id given to us?
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

	// Check if the productType exist
	productType := products.ProductType{}
	err = server.database.Product.Debug().Model(products.ProductType{}).Where("id = ?", id).Take(&productType).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	_, err = productType.DeleteProductType(server.database.Product, id)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", id))
	responses.JSON(w, http.StatusNoContent, "")
}
