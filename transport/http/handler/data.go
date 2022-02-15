package handler

import (
	"github.com/andreicalinciuc/mock-api/repository"
	"github.com/andreicalinciuc/mock-api/service"
	"github.com/andreicalinciuc/mock-api/transport/http/request"
	"github.com/andreicalinciuc/mock-api/transport/http/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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
	router.Get("/file/{path:[^/].+}", handler.GetFile)
	router.Get("/id/{id}/{path:[^/].+}", handler.GetById)
	router.Post("/{path:[^/].+}", handler.Create)
	//trebuie sa ma gandesc la o varianta de a pune parametrul dupa si de a cauta dupa el
	//
	//ca o solutie am gasit sa modific regex sa se opreasca la o anumita valoare
	//ex sa se opreasca la param si sa fac requestul: /data/test/param/{paramName}/{paramValue}
	router.Put("/id/{id}/{path:[^/].+}", handler.Update)
	router.Delete("/{path:[^/].+}", handler.Delete)
}

func (h *data) Create(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	repo := repository.NewData(params["path"])
	payload, err := request.DataArrayFromPayload(r.Body)
	if err != nil {
		return response.NewError(w, http.StatusBadRequest, "Malformed request")
	}

	err = repo.Create(nil, payload)
	if err != nil {
		return response.NewError(w, http.StatusBadRequest, err.Error())
	}

	return response.NewError(w, http.StatusCreated, "Succes ")
}

func (h *data) GetFile(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	repo := repository.NewData(params["path"])
	dataFile, err := repo.GetFile(nil)
	if err != nil {
		return response.NewError(w, http.StatusBadRequest, err.Error())
	}

	return response.New(w, http.StatusOK, dataFile)
}

func (h *data) GetById(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	repo := repository.NewData(params["path"])

	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		return response.NewError(w, http.StatusBadRequest, err.Error())
	}

	dataFind, err := repo.GetById(nil, id)
	if err != nil {
		return response.NewError(w, http.StatusBadRequest, err.Error())
	}

	return response.New(w, http.StatusOK, dataFind)

}

func (h *data) Update(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	repo := repository.NewData(params["path"])
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		return response.NewError(w, http.StatusInternalServerError, err.Error())
	}

	payloadFetch, err := request.DataFromPayload(r.Body)
	if err != nil {
		return response.NewError(w, http.StatusBadRequest, "Malformed request")
	}
	payloadFetch.Id = id
	err = repo.Update(nil, payloadFetch)
	if err != nil {
		return response.NewError(w, http.StatusBadRequest, err.Error())
	}

	return response.New(w, http.StatusOK, "Succes update")
}

func (h *data) Delete(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	repo := repository.NewData(params["path"])
	err := repo.Delete(nil)
	if err != nil {
		return response.NewError(w, http.StatusBadRequest, err.Error())
	}

	return response.New(w, http.StatusOK, "Succes delete")
}
