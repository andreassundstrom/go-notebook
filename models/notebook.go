package models

type Notebook struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}
