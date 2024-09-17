package helper

import (
	"net/http"

	"github.com/omidhaqi/clean-web-api/pkg/service_errors"
)

var StatusCodeMapping = map[string]int{
	service_errors.OtpExistsError:   409,
	service_errors.OtpUsedError:     409,
	service_errors.OtpNotValidError: 409,
}

func TranslateErrToStatusCode(err error) int {

	value , ok := StatusCodeMapping[err.Error()]

	if !ok {

		return http.StatusInternalServerError
		
	}
	return value
}
