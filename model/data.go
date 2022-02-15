package model

type Data struct {
	Id      int64                  `json:"id"`
	Payload map[string]interface{} `json:"payload"`
}
