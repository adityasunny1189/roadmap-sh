package http

import (
	"encoding/json"
	"net/http"

	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/dtos"
)

func SendJsonResponse(w http.ResponseWriter,
	status int,
	data interface{},
	errorHandlingDetails interface{}) {
	apiResponse := dtos.ApiResponse{
		StatusCode:           status,
		Data:                 data,
		ErrorHandlingDetails: errorHandlingDetails,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apiResponse)
}

func SendErrorResponse(w http.ResponseWriter, errorType dtos.ErrorType, errorMsg string) {
	var errorHandlingDetails dtos.ErrorHandlingDetails
	var errorStatusCode int

	switch errorType {
	case dtos.POPUP:
		errorStatusCode = 404
		errorHandlingDetails = dtos.ErrorHandlingDetails{
			ErrorType: dtos.POPUP,
			ErrorDetails: dtos.PopUpErrorDetails{
				Message: errorMsg,
			},
		}
	case dtos.TOAST:
		errorStatusCode = 404
		errorHandlingDetails = dtos.ErrorHandlingDetails{
			ErrorType: dtos.TOAST,
			ErrorDetails: dtos.ToastErrorDetails{
				ToastTimeoutInSec: 5,
				ToastColor:        "RED",
				Message:           errorMsg,
			},
		}
	case dtos.IN_SCREEN:
		errorStatusCode = 404
		errorHandlingDetails = dtos.ErrorHandlingDetails{
			ErrorType: dtos.IN_SCREEN,
			ErrorDetails: dtos.InScreenErrorDetails{
				Message: errorMsg,
			},
		}
	case dtos.SNACK_BAR:
		errorStatusCode = 404
		errorHandlingDetails = dtos.ErrorHandlingDetails{
			ErrorType: dtos.SNACK_BAR,
			ErrorDetails: dtos.SnakBarErrorDetails{
				Message: errorMsg,
			},
		}
	case dtos.BOTTOM_OVERLAY:
		errorStatusCode = 404
		errorHandlingDetails = dtos.ErrorHandlingDetails{
			ErrorType: dtos.BOTTOM_OVERLAY,
			ErrorDetails: dtos.BottomOverlayErrorDetails{
				IsFullScreen: false,
				ImageURL:     "",
				Heading:      "Something went wrong",
				Content:      errorMsg,
			},
		}
	}

	SendJsonResponse(w, errorStatusCode, nil, errorHandlingDetails)
}
