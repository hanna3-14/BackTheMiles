package messages

import "github.com/hanna3-14/BackTheMiles/pkg/models"

func ResultsMessage() models.ApiResponse {
	return models.ApiResponse{
		Text: "These are the results.",
	}
}
