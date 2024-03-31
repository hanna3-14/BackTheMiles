package data

import (
	"github.com/hanna3-14/BackTheMiles/pkg/models"
)

func calculateRelativePlaces(result models.Result) models.Result {
	result.RelativePlaces.Total = (result.Place.Total * 100 / result.Finisher.Total)
	result.RelativePlaces.Category = (result.Place.Category * 100 / result.Finisher.Category)
	result.RelativePlaces.Agegroup = (result.Place.Agegroup * 100 / result.Finisher.Agegroup)
	return result
}
