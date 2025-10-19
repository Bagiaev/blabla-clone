package service

import (
	"blabla-clone-api/internal/models"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Service) Register(c echo.Context) error {
	var req models.AuthRequest
	err := c.Bind(&req)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	if err := c.Validate(req); err != nil {
		s.logger.Error("Validation failed:", err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.authRepo
	existingUser, err := repo.RGetUserByEmail(req.Email)
	if err != nil && err != sql.ErrNoRows {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	if existingUser != nil {
		s.logger.Error("User already register:", existingUser)
		return c.JSON(http.StatusConflict, map[string]string{"error": "user alreadu exists"})
	}

	user := &models.User{
		Email:      req.Email,
		First_name: req.First_name,
		Last_name:  req.Last_name,
		Phone:      req.Phone,
	}
	if err := user.HashPassword(req.Password); err != nil {
		s.logger.Error("failed to hash password", err)
		return c.JSON(s.NewError(InternalServerError))
	}
	err = repo.RRegisterUser(user)
	if err != nil {
		s.logger.Error("failed to create user:", err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id":         user.ID,
		"email":      user.Email,
		"first_name": user.First_name,
		"last_name":  user.Last_name,
		"phone":      user.Phone,
	})
}

func (s *Service) Login(c echo.Context) error {
	var req models.LoginRequest
	err := c.Bind(&req)
	if err != nil {
		s.logger.Error("failed to bind request:", err)
		return c.JSON(s.NewError(InvalidParams))
	}
	err = c.Validate(req)
	if err != nil {
		s.logger.Error("validation failed:", err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.authRepo
	user, err := repo.RGetUserByEmail(req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			s.logger.Error("user not found:", req.Email)
			return c.JSON(s.NewError(InternalServerError))
		}
		s.logger.Error("database error", err)
		return c.JSON(s.NewError(InternalServerError))
	}

	if !user.ChekPassword(req.Password) {
		s.logger.Error("invalid password: ", user.Email)
		return c.JSON(s.NewError(InvalidParams))
	}

	token, err := s.jwt.GenerateToken(int(user.ID))
	if err != nil {
		s.logger.Error("failed to generate token:", err)
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"token":  token,
		"user": map[string]interface{}{
			"id":    user.ID,
			"email": user.Email,
		},
	})
}

func (s *Service) ProfileHandler(c echo.Context) error {
	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid user ID"})
	}
	repo := s.authRepo
	user, err := repo.RGetUserByID(uint(userID))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Profile data",
		"user": map[string]interface{}{
			"id":          userID,
			"first_name":  user.First_name,
			"last_name":   user.Last_name,
			"phone":       user.Phone,
			"email":       user.Email,
			"description": user.Description,
			"rating":      user.Rating,
			"created_at":  user.CreatedAt,
			"updated_at":  user.UpdatedAt,
		},
	})
}

func (s *Service) UpdateUser(c echo.Context) error { //Не меняется updated_at - нужно исправить
	var updUser struct {
		First_name  string `json:"first_name" validate:"required"`
		Last_name   string `json:"last_name" validate:"required"`
		Phone       string `json:"phone" validate:"required,phone"`
		Description string `json:"description"`
	}

	err := c.Bind(&updUser)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}
	if err := c.Validate(updUser); err != nil {
		s.logger.Error("Validation failed ", err)
		return c.JSON(s.NewError(InvalidParams))
	}
	userID, ok := c.Get("userID").(int)
	if !ok {
		s.logger.Error("userID get failed")
		return c.JSON(s.NewError(InternalServerError))
	}

	repo := s.authRepo
	user, err := repo.RGetUserByID(uint(userID))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	user.First_name = updUser.First_name
	user.Last_name = updUser.Last_name
	user.Phone = updUser.Phone
	user.Description = updUser.Description

	err = repo.RUpdateRider(user)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user": map[string]interface{}{
			"id":          user.ID,
			"first_name":  user.First_name,
			"last_name":   user.Last_name,
			"phone":       user.Phone,
			"email":       user.Email,
			"description": user.Description,
		},
	})
}

func (s *Service) UpdatePassword(c echo.Context) error {
	var updPasswordReq models.UpdPasswordRequest
	err := c.Bind(&updPasswordReq)
	if err != nil {
		s.logger.Error("failed to bind request:", err)
		return c.JSON(s.NewError(InvalidParams))
	}
	err = c.Validate(updPasswordReq)
	if err != nil {
		s.logger.Error("validation failed:", err)
		return c.JSON(s.NewError(InvalidParams))
	}

	userID, ok := c.Get("userID").(int)
	if !ok {
		return c.JSON(s.NewError(InternalServerError))
	}

	repo := s.authRepo
	user, err := repo.RGetUserByID(uint(userID))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	if !user.ChekPassword(updPasswordReq.OldPassword) {
		s.logger.Error("invalid old password", user.Email)
		return c.JSON(s.NewError(InvalidParams))
	}

	if err := user.HashPassword(updPasswordReq.NewPassword); err != nil {
		s.logger.Error("failed to hash password", err)
		return c.JSON(s.NewError(InternalServerError))
	}

	err = repo.RUpdatePassword(user)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user": map[string]interface{}{
			"email":        user.Email,
			"old_password": updPasswordReq.OldPassword,
			"new_password": updPasswordReq.NewPassword,
			"message":      "password updated successfully",
		},
	})
}
