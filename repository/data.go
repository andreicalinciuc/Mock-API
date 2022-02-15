package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/andreicalinciuc/mock-api/model"
	"io/ioutil"
	"os"
)

type dataRepo struct {
	modelPath string
}

const dataPath = "data/"
const jsonFormat = ".json"

func NewData(model string) *dataRepo {
	_ = os.Mkdir(dataPath, 0755)

	return &dataRepo{
		modelPath: dataPath + model + jsonFormat,
	}
}

func (r *dataRepo) Create(_ context.Context, data []model.Data) error {
	var dataFile []model.Data

	_, err := os.Stat(r.modelPath)
	// when err is not nil, file don't exists
	if err != nil {
		file, err := os.Create(r.modelPath)
		defer file.Close()
		if err != nil {
			return err
		}

		payloadString, err := json.Marshal(data)
		if err != nil {
			return err
		}

		_, err = file.Write(payloadString)
		if err != nil {
			return err
		}

		return nil
	}

	file, err := os.ReadFile(r.modelPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &dataFile)
	if err != nil {
		return err
	}

	for _, itemFile := range dataFile {
		for _, item := range data {
			if item.Id == itemFile.Id {
				return errors.New(fmt.Sprintf("duplicate id %v", item.Id))
			}
		}
	}

	dataFile = append(dataFile, data...)
	payloadString, err := json.Marshal(dataFile)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.modelPath, payloadString, 0755)
	if err != nil {
		return err
	}

	return nil
}

func (r *dataRepo) Update(_ context.Context, data model.Data) error {
	file, err := os.ReadFile(r.modelPath)
	if err != nil {
		return err
	}

	var dataFile []model.Data
	err = json.Unmarshal(file, &dataFile)
	if err != nil {
		return err
	}

	for index, payload := range dataFile {
		if payload.Id == data.Id {
			dataFile[index].Payload = data.Payload
			break
		}
	}

	payloadString, err := json.Marshal(dataFile)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.modelPath, payloadString, 0)
	if err != nil {
		return err
	}

	return nil
}

func (r *dataRepo) Delete(_ context.Context) error {
	err := os.Remove(r.modelPath)
	if err != nil {
		return err
	}

	return nil
}

func (r *dataRepo) GetById(_ context.Context, id int64) (model.Data, error) {
	var dataFile []model.Data
	file, err := os.ReadFile(r.modelPath)
	if err != nil {
		return model.Data{}, err
	}

	err = json.Unmarshal(file, &dataFile)
	if err != nil {
		return model.Data{}, err
	}

	for _, payload := range dataFile {
		if payload.Id == id {
			return payload, err
		}
	}

	return model.Data{}, errors.New("Invalid id")

}

func (r *dataRepo) GetFile(_ context.Context) ([]model.Data, error) {
	var dataFile []model.Data

	file, err := os.ReadFile(r.modelPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &dataFile)
	if err != nil {
		return nil, err
	}

	return dataFile, nil
}
