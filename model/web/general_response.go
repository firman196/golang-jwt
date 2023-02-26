package web

type GeneralResponse struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json: "data"`
}
