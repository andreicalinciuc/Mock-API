package handler

import (
	"fmt"
	"github.com/andreicalinciuc/mock-api/service"
	"github.com/andreicalinciuc/mock-api/transport/http/response"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

// user  godoc
type data struct {
	svc service.Data
	log service.Logger
}

// NewUser godoc
func NewUser(router service.Router, logger service.Logger) {
	handler := data{
		log: logger,
	}
	// Users Management
	router.Get("/{path:[^/].+}", handler.Get)
	router.Post("/{path:[^/].+}", handler.Create)
	router.Put("/{path:[^/].+}", handler.Update)
	router.Delete("/{path:[^/].+}", handler.Delete)
}

func (h *data) Create(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	path := params["path"]
	if checkIfFileExists(path) == false {
		file, e := os.Create("data/" + path)
		if e != nil {
			return response.NewError(w, http.StatusInternalServerError, e.Error())
		}
		file.Close()
	} else {
		return response.NewError(w, http.StatusBadRequest, "this file exists")
	}

	return response.New(w, http.StatusOK, nil)
}

func (h *data) Get(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	path := params["path"]
	fmt.Println(path)

	return response.New(w, http.StatusOK, nil)
}

func (h *data) Update(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	path := params["path"]
	fmt.Println(path)
	if checkIfFileExists(path) == false {
		fmt.Println("no exist")
	}

	return response.New(w, http.StatusOK, nil)
}

func (h *data) Delete(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	path := params["path"]
	fmt.Println(path)
	if checkIfFileExists(path) == false {
		fmt.Println("no exist")
	}

	return response.New(w, http.StatusOK, nil)

}
