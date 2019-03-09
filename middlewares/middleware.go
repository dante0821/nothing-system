package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)
func KeyAuth() echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:Session",
		Validator: func(key string, c echo.Context) (bool, error) {
			if key == "1" {
				return true, nil
			}
			return false, nil
		},
	})
}
