package domain

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Error struct {
	FailedField string
	Tag         string
	Value       string
}
