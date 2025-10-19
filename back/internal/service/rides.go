package service

import (
	"blabla-clone-api/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Функция для получения данных о поездке по id
func (s *Service) GetRideByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.ridesRepo
	ride, err := repo.RGetRideByID(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, Response{Object: ride})
}

// функция для создания новой поездки
func (s *Service) CreateRide(c echo.Context) error {
	var ride models.Ride
	err := c.Bind(&ride)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.ridesRepo
	err = repo.RCreateRide(&ride)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.String(http.StatusOK, "Ok")
}

// Функция для обновления поездки
func (s *Service) UpdateRide(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	var ride models.Ride
	if err := c.Bind(&ride); err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}
	repo := s.ridesRepo
	err = repo.RUpdateRide(id, &ride)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.String(http.StatusOK, "Ok")
}

// Функция для  удаления поездок по id
func (s *Service) DeleteRide(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}
	repo := s.ridesRepo
	err = repo.RDeleteRide(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.String(http.StatusOK, "Ok")
}

// функция для получения всех поездок
func (s *Service) GetRides(c echo.Context) error {
	repo := s.ridesRepo
	rides, err := repo.RGetRides()
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.JSON(http.StatusOK, Response{Object: rides})
}
