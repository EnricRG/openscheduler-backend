package api

import "github.com/labstack/echo/v4"

type ApiEntry struct {
	method  string
	route   string
	handler func(ctx echo.Context) error
}
