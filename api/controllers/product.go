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

func (server *Server) Products(w http.ResponseWriter, r *http.Request) {

	post := products.Product{}

	posts, err := post.AllProducts(server.database.Product)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}

func (server *Server) CreateProduct(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	product := products.Product{}
	err = json.Unmarshal(body, &product)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	product.PrepareProduct()
	err = product.ValidateProduct()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// uid, err := auth.ExtractTokenID(r)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }
	productCreated, err := product.CreateProduct(server.database.Product)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, productCreated.ID))
	responses.JSON(w, http.StatusCreated, productCreated)
}

func (server *Server) UpdateProduct(w http.ResponseWriter, r *http.Request) {

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
	product := products.Product{}
	err = server.database.Product.Debug().Model(products.Product{}).Where("id = ?", id).Take(&product).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("product not found"))
		return
	}

	// Read the data producted
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	productUpdate := products.Product{}
	err = json.Unmarshal(body, &productUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	productUpdate.PrepareProduct()
	err = productUpdate.ValidateProduct()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	productUpdated, err := productUpdate.UpdateProduct(server.database.Product, id)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, productUpdated)
}

func (server *Server) GetProductById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id64, err := strconv.ParseUint(vars["id"], 10, 64)
	id := uint32(id64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	product := products.Product{}

	productReceived, err := product.ProductById(server.database.Product, id)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, productReceived)
}

func (server *Server) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid product id given to us?
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

	// Check if the product exist
	product := products.Product{}
	err = server.database.Product.Debug().Model(products.Product{}).Where("id = ?", id).Take(&product).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	_, err = product.DeleteProduct(server.database.Product, id)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", id))
	responses.JSON(w, http.StatusNoContent, "")
}
