package model

type Response struct {
	Status  int         `json:"Status"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
}

type Error struct {
	FailedField string `json:"FailedField"`
	Tag         string `json:"Tag"`
	Value       string `json:"Value"`
}
