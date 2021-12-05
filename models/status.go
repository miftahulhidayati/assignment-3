package models

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}
type StatusDescription struct {
	Water string `json:"water"`
	Wind  string `json:"wind"`
}
type StatusShow struct {
	Status            Status            `json:"status"`
	StatusDescription StatusDescription `json:"status_description"`
}