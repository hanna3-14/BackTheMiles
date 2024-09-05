package domain

type Result struct {
	ResultID       int             `json:"resultId"`
	EventID        int             `json:"eventId"`
	Date           string          `json:"date"`
	DistanceID     int             `json:"distanceId"`
	TimeGross      RaceTime        `json:"timeGross"`
	TimeNet        RaceTime        `json:"timeNet"`
	Category       string          `json:"category"`
	Agegroup       string          `json:"agegroup"`
	Place          CategoryNumbers `json:"place"`
	Finisher       CategoryNumbers `json:"finisher"`
	RelativePlaces CategoryNumbers `json:"relativePlaces"`
}

type RaceTime struct {
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
	Seconds int `json:"seconds"`
}

type CategoryNumbers struct {
	Total    int `json:"total"`
	Category int `json:"category"`
	Agegroup int `json:"agegroup"`
}
