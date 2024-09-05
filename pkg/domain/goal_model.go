package domain

type Goal struct {
	ID       int    `json:"id"`
	Distance string `json:"distance"`
	Time     string `json:"time"`
}
