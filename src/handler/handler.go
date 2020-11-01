package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Hello localhost:8090でHello, Worldを表示
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// HelloUser Path Paramtersの利用
func HelloUser(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, "Hello, "+name+"!!")
}

// UserInfo JSONの返却サンプル
func UserInfo(c echo.Context) error {
	user := user{
		Name:     "kohei.sato",
		Age:      42,
		Position: "Engineer",
	}
	return c.JSON(http.StatusOK, user)
}

type user struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Position string `json:"position"`
}
