package messages

import "github.com/hanna3-14/BackTheMiles/pkg/models"

func PublicMessage() models.ApiResponse {
	return models.ApiResponse{
		Text: "This is a public message.",
	}
}

func ProtectedMessage() models.ApiResponse {
	return models.ApiResponse{
		Text: "This is a protected message.",
	}
}

func AdminMessage() models.ApiResponse {
	return models.ApiResponse{
		Text: "This is an admin message.",
	}
}

func ResultsMessage() models.ApiResponse {
	return models.ApiResponse{
		Text: "These are the results.",
	}
}
