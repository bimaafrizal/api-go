package helper

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ResponseError struct {
	Code    int    `bson:"code" json:"code"`
	Message string `bson:"message" json:"message"`
}

func InternalServerErrorResponse(msg string) error {
	errResponse := ResponseError{Code: http.StatusInternalServerError, Message: msg}
	return echo.NewHTTPError(http.StatusInternalServerError, errResponse)
}
func NotFoundErrorResponse() error {
	errResponse := ResponseError{Code: http.StatusNotFound, Message: "Data Not Found"}
	return echo.NewHTTPError(http.StatusNotFound, errResponse)
}
func BadRequestResponse(msg string) error {
	errResponse := ResponseError{Code: http.StatusBadRequest, Message: msg}
	return echo.NewHTTPError(http.StatusBadRequest, errResponse)
}
