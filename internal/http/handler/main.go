package handler

import "github.com/labstack/echo/v4"

type Chat interface {
	GetMessage(ctx echo.Context) error
	SendMessage(ctx echo.Context) error
}

func GetMessage(ctx echo.Context) error {
	return nil
}

func SendMessage(ctx echo.Context) error {
	return nil
}
