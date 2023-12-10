package utils

type ApiResponseWithData[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message"`
}

type ApiResponseOnlyMessage struct {
	Message string `json:"message"`
}
