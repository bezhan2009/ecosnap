package controllers

import "ecosnap/internal/app/models"

func newErrorResponse(error string) models.ErrorResponse {
	return models.ErrorResponse{
		Error: error,
	}
}
