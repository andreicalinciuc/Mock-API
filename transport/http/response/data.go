package response

import "github.com/andreicalinciuc/mock-api/model"

type Data struct {
	ID   int64  `json:"id"`
	Body string `json:"body"`
}

// FromUser godoc
func FromUser(user *model.Data) Data {

	return Data{}
}

// FromUsersModel godoc
func FromUsersModel(users []model.Data, count uint64, err error) ([]Data, uint64, error) {
	if err != nil {
		return nil, count, err
	}
	result := make([]Data, len(users))
	for i, user := range users {
		result[i] = FromUser(&user)
	}
	return result, count, err
}
