package handlerMessage

import (
	"net/http"

	"github.com/pkg/errors"
)

var (
	// ErrRetrieveData in case an error during retrieve data from DB
	ErrRetrieveData = errors.New("Error during retrieve Data")
	// ErrFindData in case an error during find data from DB
	ErrFindData = errors.New("Error during find Data")
	// ErrConnection in case an error during connection with DB
	ErrConnection = errors.New("Error during connection with DB")
	// ErrCheckConnection in case an error during connection with DB
	ErrCheckConnection = errors.New("Error during check connection with DB")
	// ErrDataNotFound in case an error during connection with DB
	ErrDataNotFound = errors.New("No data available")
	// ErrDataAlreadyPresent in case the key already exist into DB
	ErrDataAlreadyPresent = errors.New("Data already present into DB")
)

// translate the internal error code to http StatusCode and related message/code
func ToStatusCodeMessage(err error) (int, string, int) {
	switch err {
	case ErrRetrieveData:
		return http.StatusBadRequest, err.Error(), 1
	case ErrFindData:
		return http.StatusInternalServerError, err.Error(), 2
	case ErrConnection:
		return http.StatusInternalServerError, err.Error(), 3
	case ErrCheckConnection:
		return http.StatusInternalServerError, err.Error(), 4
	case ErrDataNotFound:
		return http.StatusNotFound, err.Error(), 5
	case ErrDataAlreadyPresent:
		return http.StatusConflict, err.Error(), 6
	default:
		return http.StatusInternalServerError, "Internal Error", 7
	}
}
