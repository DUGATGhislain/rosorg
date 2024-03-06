package middleware

import (
	"github.com/labstack/echo/v4"
)

func LocalizeApp() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			locale := c.Get("locale")
			if locale == "" || locale == nil {
				locale = c.Request().Header.Get("Accept-Language")
			}
			if locale == "" || locale == nil {
				locale = "en-US"
			}
			c.Set("locale", locale)
			return next(c)
		}
	}
}
