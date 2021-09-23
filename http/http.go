package http

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Success bool        `json:"success"`
	ErrMsg  string      `json:"error,omitempty"`
}

func getIDFromRequest(strID string) (int, error) {
	errText := "Укажите ID"
	if strID == "" {
		return 0, errors.New(errText)
	}
	ID, err := strconv.Atoi(strID)
	if err != nil {
		return 0, errors.New(errText)
	}
	return ID, nil
}

func respondWithErrorStatus(c echo.Context, status int, errMsg string) error {
	return c.JSON(status, &Response{Success: false, ErrMsg: errMsg})
}

func respondWithData(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, &Response{Success: true, Data: data})
}

func respondWithNoData(c echo.Context) error {
	return c.JSON(http.StatusOK, &Response{Success: true})
}
