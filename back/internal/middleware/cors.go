package middleware

import "github.com/labstack/echo/v4"

func CORS() echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            c.Response().Header().Set("Access-Control-Allow-Origin", "*")
            c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH, HEAD")
            c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept, Origin, Access-Control-Allow-Headers, Access-Control-Request-Method, Access-Control-Request-Headers")
            c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
            c.Response().Header().Set("Access-Control-Max-Age", "86400")
            
            if c.Request().Method == "OPTIONS" {
                return c.NoContent(200)
            }
            
            return next(c)
        }
    }
}