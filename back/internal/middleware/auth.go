package middleware

import (
	"blabla-clone-api/pkg/jwt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func JWTMiddleware(j *jwt.JWT) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing token"})
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := j.ParseToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
			}

			// Преобразуем UserID в int
			userID, err := strconv.Atoi(claims.UserID)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid user ID in token"})
			}

			c.Set("userID", userID)
			return next(c)
		}
	}
}
