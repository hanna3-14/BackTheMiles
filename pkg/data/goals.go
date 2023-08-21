package data

import "github.com/hanna3-14/BackTheMiles/pkg/models"

func GoalsData() []models.Goal {
	return []models.Goal{
		{
			ID:       1,
			Distance: "M",
			Time:     "04:59:59",
		},
		{
			ID:       2,
			Distance: "HM",
			Time:     "01:59:59",
		},
	}
}
