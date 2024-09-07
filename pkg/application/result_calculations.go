package application

import "github.com/hanna3-14/BackTheMiles/pkg/domain"

func calculateRelativePlaces(result domain.Result) domain.Result {
	result.RelativePlaces.Total = (result.Place.Total * 100 / result.Finisher.Total)
	result.RelativePlaces.Category = (result.Place.Category * 100 / result.Finisher.Category)
	result.RelativePlaces.Agegroup = (result.Place.Agegroup * 100 / result.Finisher.Agegroup)
	return result
}
