package web

type GeneralResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json: "data"`
}
