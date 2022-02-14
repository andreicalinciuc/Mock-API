package handler

import (
	"encoding/json"
	"github.com/andreicalinciuc/mock-api/model"
	"github.com/andreicalinciuc/mock-api/service"
	"github.com/andreicalinciuc/mock-api/transport/http/request"
	"github.com/andreicalinciuc/mock-api/transport/http/response"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
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
	router.Put("/id/{id}/{path:[^/].+}", handler.Update)
	router.Delete("/{path:[^/].+}", handler.Delete)
}

const dataPath = "data/"

func (h *data) Create(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	path := dataPath + params["path"]
	payload, err := request.DataArrayFromPayload(r.Body)
	if err != nil {
		return response.NewError(w, http.StatusBadRequest, "Malformed request")
	}

	switch r.Header.Get("Type") {
	case "file":
		{
			if checkIfFileExists(path) == false {
				file, err := os.Create(path)
				defer file.Close()
				if err != nil {
					return response.NewError(w, http.StatusInternalServerError, err.Error())
				}

				payloadString, err := json.Marshal(payload)
				if err != nil {
					return response.NewError(w, http.StatusInternalServerError, err.Error())
				}

				_, err = file.Write(payloadString)
				if err != nil {
					return response.NewError(w, http.StatusInternalServerError, err.Error())
				}

			} else {
				return response.NewError(w, http.StatusBadRequest, "this file exists")
			}

			break
		}
	case "payload":
		{
			file, err := os.ReadFile(path)
			if err != nil {
				return response.NewError(w, http.StatusInternalServerError, err.Error())
			}

			var dataFile []model.Data
			err = json.Unmarshal(file, &dataFile)
			if err != nil {
				return response.NewError(w, http.StatusInternalServerError, err.Error())
			}

			for _, item := range dataFile {
				for _, data := range payload {
					if data.Id == item.Id {
						return response.NewError(w, http.StatusBadRequest, "duplicate id "+strconv.Itoa(int(data.Id)))
					}
				}
			}

			dataFile = append(dataFile, payload...)
			payloadString, err := json.Marshal(dataFile)
			if err != nil {
				return response.NewError(w, http.StatusInternalServerError, err.Error())
			}

			err = ioutil.WriteFile(path, payloadString, 0)
			if err != nil {
				return response.NewError(w, http.StatusInternalServerError, err.Error())
			}
		}

	default:
		{
			return response.NewError(w, http.StatusBadRequest, "unknown type action")
		}
	}
	return response.NewError(w, http.StatusOK, "Succes ")
}

func (h *data) GetFile(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	path := dataPath + params["path"]
	var dataFile []model.Data

	file, err := os.ReadFile(path)
	if err != nil {
		return response.NewError(w, http.StatusInternalServerError, err.Error())
	}

	err = json.Unmarshal(file, &dataFile)
	if err != nil {
		return response.NewError(w, http.StatusInternalServerError, err.Error())
	}

	return response.New(w, http.StatusOK, dataFile)
}

func (h *data) GetById(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	path := dataPath + params["path"]
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		return response.NewError(w, http.StatusBadRequest, err.Error())
	}

	var dataFile []model.Data
	file, err := os.ReadFile(path)
	if err != nil {
		return response.NewError(w, http.StatusBadRequest, err.Error())
	}

	err = json.Unmarshal(file, &dataFile)
	if err != nil {
		return response.NewError(w, http.StatusInternalServerError, err.Error())
	}

	for _, payload := range dataFile {
		if payload.Id == id {
			return response.New(w, http.StatusOK, payload)
		}
	}

	return response.NewError(w, http.StatusNotFound, "invalid id")
}

func (h *data) Update(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	path := dataPath + params["path"]
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		return response.NewError(w, http.StatusInternalServerError, err.Error())
	}

	payloadFetch, err := request.DataFromPayload(r.Body)
	if err != nil {
		return response.NewError(w, http.StatusBadRequest, "Malformed request")
	}

	file, err := os.ReadFile(path)
	if err != nil {
		return response.NewError(w, http.StatusInternalServerError, err.Error())
	}

	var dataFile []model.Data
	err = json.Unmarshal(file, &dataFile)
	if err != nil {
		return response.NewError(w, http.StatusInternalServerError, err.Error())
	}

	for index, payload := range dataFile {
		if payload.Id == id {
			dataFile[index].Payload = payloadFetch.Payload
			break
		}
	}

	payloadString, err := json.Marshal(dataFile)
	if err != nil {
		return response.NewError(w, http.StatusInternalServerError, err.Error())
	}

	err = ioutil.WriteFile(path, payloadString, 0)
	if err != nil {
		return response.NewError(w, http.StatusInternalServerError, err.Error())
	}

	return response.New(w, http.StatusOK, "Succes update")
}

func (h *data) Delete(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	path := dataPath + params["path"]
	err := os.Remove(path)
	if err != nil {
		return response.NewError(w, http.StatusBadRequest, err.Error())
	}

	return response.New(w, http.StatusOK, "Succes delete")

}
