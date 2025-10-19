package service

import (
	"blabla-clone-api/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Функция для создания отзыва
func (s *Service) CreateReport(c echo.Context) error {
	var reportReq models.Report
	err := c.Bind(&reportReq)
	if err != nil {
		s.logger.Error("Bind (title, description) failed: ", err)
		return c.JSON(s.NewError(InvalidParams))
	}
	if err = c.Validate(reportReq); err != nil {
		s.logger.Error("Validation failed: ", err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.reportRepo
	err = repo.RCreateReport(&reportReq)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id":          reportReq.ID,
		"title":       reportReq.Title,
		"description": reportReq.Description,
		"created_at":  reportReq.CreatedAt,
	})
}

// Функция для получения отзыва по id
func (s *Service) GetReportByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.reportRepo
	ride, err := repo.RGetReportByID(uint(id))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, Response{Object: ride})
}

// Функция для получения всех отзывов пользователя
func (s *Service) GetRiderReports(c echo.Context) error {
	userID := c.Get("userID").(int)

	repo := s.reportRepo
	reports, err := repo.RGetReports(uint(userID))
	if err != nil {
		s.logger.Error("")
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.JSON(http.StatusOK, Response{Object: reports})
}

//Функция для удаления отзыва
