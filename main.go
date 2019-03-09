package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"time"
	"net/http"
	"MySystem/middlewares"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"*"},
			AllowCredentials: true,
			MaxAge:           int(time.Hour) * 12,
		},
	))
	route(e)
	e.Start(":8111")
}

func route(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status},form=${form}\n",
	}))
	//group := e.Group("/")
	e.POST("login", Login)
	e.GET("user_info", GetCurrentUserInfo, middlewares.KeyAuth())
	e.File("/", "web/html/login.html")
	e.File("/login", "web/html/login.html")
	e.File("/index", "web/html/index.html")
	e.File("/canvas", "web/html/canvas.html")
	e.Static("/src", "web/src")
	e.Static("/lay", "web/src/lay")
	e.Static("/css", "web/src/css")
	e.Static("/html", "web/html")
}

func Login(c echo.Context) error {
	loginInfo := &LoginInfo{}
	if err:= c.Bind(loginInfo); err != nil {
		return c.JSON(http.StatusBadRequest, &Rsp{"请求体错误"})
	}

	if loginInfo.Username != "szc" {
		return c.JSON(http.StatusBadRequest, &Rsp{"账号错误"})
	}
	if loginInfo.Password != "szc" {
		return c.JSON(http.StatusBadRequest, &Rsp{"密码错误"})
	}
	return c.JSON(http.StatusOK, nil)
}

type Rsp struct {
	Msg string `json:"msg"`
}

type LoginInfo struct {
	Username string
	Password string
}

func GetCurrentUserInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}