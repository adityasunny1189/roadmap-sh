package dtos

type ErrorType string

const (
	BOTTOM_OVERLAY ErrorType = "BOTTOM_OVERLAY"
	SNACK_BAR      ErrorType = "SNACK_BAR"
	IN_SCREEN      ErrorType = "IN_SCREEN"
	POPUP          ErrorType = "POPUP"
	TOAST          ErrorType = "TOAST"
)

type ApiResponse struct {
	StatusCode           int         `json:"status"`
	Data                 interface{} `json:"data,omitempty"`
	ErrorHandlingDetails interface{} `json:"errorHandlingDetails,omitempty"`
}

type ErrorHandlingDetails struct {
	ErrorType    ErrorType   `json:"errorType"`
	ErrorDetails interface{} `json:"errorDetails"`
}

type BottomOverlayErrorDetails struct {
	IsFullScreen bool   `json:"isFullScreen"`
	ImageURL     string `json:"imageUrl"`
	Heading      string `json:"heading"`
	Content      string `json:"content"`
}

type SnakBarErrorDetails struct {
	Message string `json:"message"`
}

type InScreenErrorDetails struct {
	Message string `json:"message"`
}

type PopUpErrorDetails struct {
	Message string `json:"message"`
}

type ToastErrorDetails struct {
	ToastTimeoutInSec int    `json:"toastTimeoutInSec"`
	ToastColor        string `json:"toastColor"`
	Message           string `json:"message"`
}
