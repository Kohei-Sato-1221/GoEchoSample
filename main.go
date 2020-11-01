package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"handler"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORSの設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8090"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/", handler.Hello)

	//Path Parameters
	//セミコロン部分がパラメータ
	e.GET("/hello/:name", handler.HelloUser)

	//JSONの返却
	e.GET("/userinfo", handler.UserInfo)

	e.Logger.Fatal(e.Start(":8090"))
}
