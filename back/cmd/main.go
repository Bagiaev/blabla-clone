package main

import (
	"blabla-clone-api/internal/middleware"
	"blabla-clone-api/internal/service"
	"blabla-clone-api/pkg/jwt"
	"blabla-clone-api/pkg/logs"
	"blabla-clone-api/pkg/validator"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	//создаем логгер
	logger := logs.NewLogger(false)

	//подключение к бд
	db, err := PostgresConnection()
	if err != nil {
		logger.Fatal(err)
	}

	//jwt пока пропускаем
	jwtSecret := jwt.NewJWT(os.Getenv("JWT_SECRET"))

	svc := service.NewService(db, logger, *jwtSecret)

	router := echo.New()
	router.Use(middleware.CORS())
	//валидатор
	router.Validator = validator.New()
	api := router.Group("/api")

	//ручки
	api.POST("/register", svc.Register)
	api.POST("/login", svc.Login)

	protected := api.Group("/user")
	protected.Use(middleware.JWTMiddleware(jwtSecret))
	//профиль
	protected.GET("/profile", svc.ProfileHandler)
	protected.PUT("/password-reset", svc.UpdatePassword)
	protected.PUT("/information-update", svc.UpdateUser)

	//Поездки
	protected.GET("/rides/:id", svc.GetRideByID)
	protected.GET("/rides", svc.GetRides)
	protected.POST("/rides", svc.CreateRide)
	protected.PUT("/rides/:id", svc.UpdateRide)
	protected.DELETE("/rides/:id", svc.DeleteRide)

	router.Logger.Fatal(router.Start("0.0.0.0:8000"))
}
