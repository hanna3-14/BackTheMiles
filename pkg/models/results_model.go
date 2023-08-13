package models

type Result struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Distance string `json:"distance"`
	Time     string `json:"time"`
	Place    int    `json:"place"`
}
