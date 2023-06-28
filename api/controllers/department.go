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

func (server *Server) Departments(w http.ResponseWriter, r *http.Request) {

	post := products.Department{}

	posts, err := post.AllDepartments(server.database.Product)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}

func (server *Server) CreateDepartment(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	department := products.Department{}
	err = json.Unmarshal(body, &department)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	department.PrepareDepartment()
	err = department.ValidateDepartment()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// uid, err := auth.ExtractTokenID(r)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }
	departmentCreated, err := department.CreateDepartment(server.database.Product)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, departmentCreated.ID))
	responses.JSON(w, http.StatusCreated, departmentCreated)
}

func (server *Server) UpdateDepartment(w http.ResponseWriter, r *http.Request) {

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
	department := products.Department{}
	err = server.database.Product.Debug().Model(products.Department{}).Where("id = ?", id).Take(&department).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("department not found"))
		return
	}

	// Read the data departmented
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	departmentUpdate := products.Department{}
	err = json.Unmarshal(body, &departmentUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	departmentUpdate.PrepareDepartment()
	err = departmentUpdate.ValidateDepartment()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	departmentUpdated, err := departmentUpdate.UpdateDepartment(server.database.Product, id)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, departmentUpdated)
}

func (server *Server) GetDepartmentById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id64, err := strconv.ParseUint(vars["id"], 10, 64)
	id := uint32(id64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	department := products.Department{}

	departmentReceived, err := department.DepartmentById(server.database.Product, id)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, departmentReceived)
}

func (server *Server) DeleteDepartment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid department id given to us?
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

	// Check if the department exist
	department := products.Department{}
	err = server.database.Product.Debug().Model(products.Department{}).Where("id = ?", id).Take(&department).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	_, err = department.DeleteDepartment(server.database.Product, id)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", id))
	responses.JSON(w, http.StatusNoContent, "")
}
