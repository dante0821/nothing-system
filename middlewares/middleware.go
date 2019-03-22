package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)
func KeyAuth() echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:Session",
		Validator: func(key string, c echo.Context) (bool, error) {
			if key == "A81E2425F5B76A419A97CED421B23852" {
				c.Set("id", 1)
				return true, nil
			}
			//DEBUG
			c.Set("id", 1)
			return true, nil
		},
	})
}
