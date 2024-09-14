package rest

import (
	"encoding/json"
	"net/http"

	"github.com/adityasunny1189/url-shortner/internal/models"
)

func sendJsonResponse(w http.ResponseWriter,
	status int,
	data interface{},
	errorHandlingDetails interface{}) {
	apiResponse := models.ApiResponse{
		StatusCode:           status,
		Data:                 data,
		ErrorHandlingDetails: errorHandlingDetails,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apiResponse)
}

func sendErrorResponse(w http.ResponseWriter, errorType models.ErrorType, errorMsg string) {
	var errorHandlingDetails models.ErrorHandlingDetails
	var errorStatusCode int

	switch errorType {
	case models.POPUP:
		errorStatusCode = 404
		errorHandlingDetails = models.ErrorHandlingDetails{
			ErrorType: models.POPUP,
			ErrorDetails: models.PopUpErrorDetails{
				Message: errorMsg,
			},
		}
	case models.TOAST:
		errorStatusCode = 404
		errorHandlingDetails = models.ErrorHandlingDetails{
			ErrorType: models.TOAST,
			ErrorDetails: models.ToastErrorDetails{
				ToastTimeoutInSec: 5,
				ToastColor:        "RED",
				Message:           errorMsg,
			},
		}
	case models.IN_SCREEN:
		errorStatusCode = 404
		errorHandlingDetails = models.ErrorHandlingDetails{
			ErrorType: models.IN_SCREEN,
			ErrorDetails: models.InScreenErrorDetails{
				Message: errorMsg,
			},
		}
	case models.SNACK_BAR:
		errorStatusCode = 404
		errorHandlingDetails = models.ErrorHandlingDetails{
			ErrorType: models.SNACK_BAR,
			ErrorDetails: models.SnakBarErrorDetails{
				Message: errorMsg,
			},
		}
	case models.BOTTOM_OVERLAY:
		errorStatusCode = 404
		errorHandlingDetails = models.ErrorHandlingDetails{
			ErrorType: models.BOTTOM_OVERLAY,
			ErrorDetails: models.BottomOverlayErrorDetails{
				IsFullScreen: false,
				ImageURL:     "",
				Heading:      "Something went wrong",
				Content:      errorMsg,
			},
		}
	}

	sendJsonResponse(w, errorStatusCode, nil, errorHandlingDetails)
}
