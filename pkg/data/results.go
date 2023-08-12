package data

import "github.com/hanna3-14/BackTheMiles/pkg/models"

func ResultsData() models.Results {
	return models.Results{
		Results: []models.Result{
			{
				ID:       1,
				Name:     "Baden-Marathon",
				Distance: "HM",
				Time:     "02:21:40",
				Place:    2336,
			},
			{
				ID:       2,
				Name:     "Schwarzwald-Marathon",
				Distance: "HM",
				Time:     "02:09:45",
				Place:    535,
			},
			{
				ID:       3,
				Name:     "Bienwald-Marathon",
				Distance: "HM",
				Time:     "02:09:14",
				Place:    928,
			},
			{
				ID:       4,
				Name:     "Freiburg-Marathon",
				Distance: "M",
				Time:     "05:29:09",
				Place:    916,
			},
		},
	}
}
