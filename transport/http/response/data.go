package response

import "github.com/andreicalinciuc/mock-api/model"

type Data struct {
	Id      int64                  `json:"id"`
	Payload map[string]interface{} `json:"payload"`
}

// FromData godoc
func FromData(data *model.Data) Data {
	return Data{
		Id:      data.Id,
		Payload: data.Payload,
	}
}

// FromDataModel godoc
func FromDataModel(datas []model.Data, count uint64, err error) ([]Data, uint64, error) {
	if err != nil {
		return nil, count, err
	}
	result := make([]Data, len(datas))
	for i, data := range datas {
		result[i] = FromData(&data)
	}
	return result, count, err
}
