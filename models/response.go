package models

// Response is a standardized response structure
type Response struct {
	Status  string      `json:"status"`  // success or error
	Message string      `json:"message"` // detailed message
	Data    interface{} `json:"data"`    // actual response data
}
