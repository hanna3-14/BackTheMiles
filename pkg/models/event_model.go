package models

type Event struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	ResultIDs []int  `json:"resultIds"`
}
