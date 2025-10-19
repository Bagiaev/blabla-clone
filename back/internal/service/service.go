package service

import (
	"blabla-clone-api/internal/repo/auth"
	"blabla-clone-api/internal/repo/reports"
	"blabla-clone-api/internal/repo/rides"
	"blabla-clone-api/pkg/jwt"
	"blabla-clone-api/pkg/validator"
	"database/sql"

	"github.com/labstack/echo/v4"
)

const (
	InvalidParams       = "invalid params"
	InternalServerError = "internal error"
)

type Service struct {
	db        *sql.DB
	logger    echo.Logger
	validator *validator.CustomValidator

	ridesRepo  *rides.Repo
	authRepo   *auth.Repo
	reportRepo *reports.Repo

	jwt *jwt.JWT
}

func NewService(db *sql.DB, logger echo.Logger, jwtSecret jwt.JWT) *Service {
	svc := &Service{
		db:        db,
		logger:    logger,
		validator: validator.New(),
		jwt:       &jwtSecret,
	}
	svc.initRepositories(db)

	return svc
}

func (s *Service) initRepositories(db *sql.DB) {
	s.ridesRepo = rides.NewRepo(db)
	s.authRepo = auth.NewRepo(db)
	s.reportRepo = reports.NewRepo(db)
}

type Response struct {
	Object       any    `json:"object,omitempty"`
	ErrorMessage string `json:"error,omitempty"`
}

func (r *Response) Error() string {
	return r.ErrorMessage
}

func (s *Service) NewError(err string) (int, *Response) {
	return 400, &Response{ErrorMessage: err}
}
