package helper

import validation "github.com/omidhaqi/clean-web-api/api/validations"

type BaseHttpResponse struct {
	Result           any                           `json:"result"`
	Success          bool                          `json:"success"`
	ResultCode       int                           `json:"resultCode"`
	ValidationErrors *[]validation.ValidationError `json:"validationErrors"`
	Error            any                           `json:"error"`
}

func GenerateBaseResponse(result any, success bool, resultCode int) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Success:    success,
		ResultCode: resultCode,
	}
}

func GenerateBaseResponseWithError(result any, success bool, resultCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Success:    success,
		ResultCode: resultCode,
		Error:      err.Error(),
	}

}

func GenerateBaseResponseWithAnyError(result any, success bool, resultCode int, err any) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Success:    success,
		ResultCode: resultCode,
		Error:      err,
	}
}

func GenerateBaseResponseWithValidationError(result any, success bool, resultCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Success:          success,
		ResultCode:       resultCode,
		ValidationErrors: validation.GetValidationErrors(err),
	}
}
