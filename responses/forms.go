package responses

import "digimovie/src/validations"

type baseResponse struct {
	Status           bool  `json:"status" binding:"required"`
	StatusCode       int   `json:"status_code" binding:"required"`
	Result           any   `json:"result"`
	Error            string `json:"error"`
	ValidationErrors []validations.ValidationError `json:"validation_error"`
}

func GenerateNormalResponse(status bool, statusCode int, result any) *baseResponse {
	return &baseResponse{
		Status: status, StatusCode: statusCode, Result: result,
	}
}

func GenerateResponseWithError(status bool, statusCode int, err error) *baseResponse {
	return &baseResponse{
		Status: status, StatusCode: statusCode, Error: err.Error(),
	}
}

func GenerateResponseWithValidationError(status bool, statusCode int, err error) *baseResponse {
	return &baseResponse{
		Status: status, StatusCode: statusCode, ValidationErrors: *validations.GetValidationErrors(err),
	}
}