package models

type Distance struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	DistanceInMeters int    `json:"distanceInMeters"`
}
