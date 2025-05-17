package services

type GetAllStatusResponseBody struct {
	Up   int `json:"up"`
	Down int `json:"down"`
}
