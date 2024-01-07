package models

type Result struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Distance         string `json:"distance"`
	TimeGross        string `json:"timeGross"`
	TimeNet          string `json:"timeNet"`
	Category         string `json:"category"`
	Agegroup         string `json:"agegroup"`
	PlaceTotal       int    `json:"placeTotal"`
	PlaceCategory    int    `json:"placeCategory"`
	PlaceAgegroup    int    `json:"placeAgegroup"`
	FinisherTotal    int    `json:"finisherTotal"`
	FinisherCategory int    `json:"finisherCategory"`
	FinisherAgegroup int    `json:"finisherAgegroup"`
}
