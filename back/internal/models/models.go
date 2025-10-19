package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	First_name   string    `json:"first_name" validate:"required"`
	Last_name    string    `json:"last_name" validate:"required"`
	Email        string    `json:"email" validate:"required,email"`
	PasswordHash string    `json:"-"`
	Phone        string    `json:"phone" validate:"required"` //validate ??
	Is_driver    bool      `json:"is_driver"`
	Rating       float32   `json:"rating"`
	Description  string    `json:"description"`
	Car_model    string    `json:"car_model"`
}
type AuthRequest struct {
	First_name string `json:"first_name" validate:"required"`
	Last_name  string `json:"last_name" validate:"required"`
	Phone      string `json:"phone" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}
type UpdPasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}

type Car struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uint      `json:"user_id"`
	Brand     string    `json:"brand" validate:"required"`
	Model     string    `json:"model" validate:"required"`
	Year      int       `json:"year"`
	Seats     int       `json:"seats"`
}

type Ride struct {
	ID             uint          `json:"id"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
	DriverID       uint          `json:"driver_id"`
	Driver         *User         `json:"driver,omitempty"`
	CarID          uint          `json:"car_id"`
	FromLocation   string        `json:"from_location" validate:"required"`
	ToLocation     string        `json:"to_location" validate:"required"`
	DepartureAt    time.Time     `json:"departure_at" validate:"required"`
	SeatsAvailable int           `json:"seats_available"`
	PriceCents     int           `json:"price_cents" validate:"required"`
	Description    string        `json:"description"`
	RideRequests   []RideRequest `json:"ride_requests,omitempty"`
}

type RideRequest struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	RideID      uint           `json:"ride_id"`
	PassengerID uint           `json:"passenger_id"`
	Passenger   *User          `gorm:"foreignKey:PassengerID" json:"passenger,omitempty"`
	Status      string         `json:"status"` // pending, accepted, rejected
}

type Notification struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    uint           `json:"user_id"`
	Type      string         `json:"type"`
	Message   string         `json:"message"`
	Read      bool           `json:"read"`
}

type Report struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
}

func (u *User) HashPassword(password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hashed)
	return nil
}

func (u *User) ChekPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}
