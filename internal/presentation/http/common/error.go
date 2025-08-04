package common

type InternalErrorDto struct {
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

func CreateError(
	message string,
	error error,
) *InternalErrorDto {
	if error != nil {
		return &InternalErrorDto{
			Message: message,
			Error:   error.Error(),
		}
	}

	return &InternalErrorDto{
		Message: message,
		Error:   error,
	}
}
