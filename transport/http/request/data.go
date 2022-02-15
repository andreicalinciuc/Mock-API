package request

import (
	"github.com/andreicalinciuc/mock-api/model"
	"io"
)

type Data struct {
	Id      int64                  `json:"id"`
	Payload map[string]interface{} `json:"payload"`
}

// ToModel Helper function to convert request.User to []model.User
func (data *Data) ToModel() model.Data {
	return model.Data{
		Id:      data.Id,
		Payload: data.Payload,
	}
}

func DataFromPayload(reader io.ReadCloser) (model.Data, error) {
	var u Data
	err := Unmarshal(reader, &u)
	if err != nil {
		return model.Data{}, err
	}
	return u.ToModel(), nil
}

func DataArrayFromPayload(reader io.ReadCloser) ([]model.Data, error) {
	//var data []Data
	var dataModel []model.Data
	err := Unmarshal(reader, &dataModel)
	if err != nil {
		return nil, err
	}

	return dataModel, nil
}
