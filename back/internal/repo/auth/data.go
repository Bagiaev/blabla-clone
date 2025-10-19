package auth

import (
	"blabla-clone-api/internal/models"
	"database/sql"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

// реализовать сброс пароля (забыли пароль). auth/refresh позволит получать новые токены без повторного логина, пароля. logout добавляет токен в черный список
// GET /users/search?name=:name поиск пользователей
// Функция для регистрации (добавления в бд пользователя)
func (r *Repo) RRegisterUser(user *models.User) error {
	err := r.db.QueryRow(
		`INSERT INTO users (first_name, last_name, phone, email, hashed_password, created_at)
		VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP)
		RETURNING id, created_at`,
		user.First_name,
		user.Last_name,
		user.Phone,
		user.Email,
		user.PasswordHash,
	).Scan(&user.ID, &user.CreatedAt)
	return err
}

// Функция для получения информации о пользователе по email
func (r *Repo) RGetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(`SELECT id, created_at, first_name, last_name, email, description, rating, hashed_password, phone FROM users WHERE email = $1`, email).
		Scan(&user.ID, &user.CreatedAt, &user.First_name, &user.Last_name, &user.Email, &user.Description, &user.Rating, &user.PasswordHash, &user.Phone)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Функция для получения информации о пользователе по id
func (r *Repo) RGetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(`SELECT id, created_at, first_name, last_name, email, description, rating, hashed_password, phone FROM users WHERE id = $1`, id).
		Scan(&user.ID, &user.CreatedAt, &user.First_name, &user.Last_name, &user.Email, &user.Description, &user.Rating, &user.PasswordHash, &user.Phone)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// users/me для получения личного профиля пользователя (БД)
func (r *Repo) RGetUserPrivate(id int) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(
		`SELECT id, created_at, first_name, last_name, email, hashed_password, description, rating, phone 
		FROM users WHERE email = $1`, id,
	).Scan(&user.ID, &user.CreatedAt, &user.First_name, &user.Last_name, &user.Email, &user.PasswordHash, &user.Description, &user.Rating, &user.Phone)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Функция для обновления данных пассажира (БД)
func (r *Repo) RUpdatePassenger(user *models.User) error {
	err := r.db.QueryRow(
		`UPDATE users SET first_name = $1, last_name = $2, phone = $3, updated_at = CURRENT_TIMESTAMP
		WHERE id = $4
		RETURNING updated_at`,
		user.First_name, user.Last_name, user.Phone, user.ID,
	).Scan(&user.UpdatedAt)
	return err
}

// Функция для обновления данных водителя (БД)
func (r *Repo) RUpdateRider(user *models.User) error {
	err := r.db.QueryRow(
		`UPDATE users SET first_name = $1, last_name = $2, description = $3, phone = $4, updated_at = CURRENT_TIMESTAMP
		WHERE id = $5
		RETURNING updated_at`,
		user.First_name, user.Last_name, user.Description, user.Phone, user.ID,
	).Scan(&user.UpdatedAt)
	return err
}

// функция для обновления пароля (БД)
func (r *Repo) RUpdatePassword(user *models.User) error {
	_, err := r.db.Exec(
		`UPDATE users SET hashed_password = $1
		WHERE id = $2`,
		user.PasswordHash,
		user.ID,
	)
	return err
}

// Функция для того чтобы пользователь стал водителем (БД)
func (r *Repo) RBecomeRider(user *models.User) error {
	_, err := r.db.Exec(
		`UPDATE users SET is_driver = TRUE, rating = 0, description = $1, car_model = $2
		WHERE id = $3`, user.Description, user.Car_model, user.ID,
	)
	if err != nil {
		return err
	}
	return nil
} // фронт: реквест для этой функции
