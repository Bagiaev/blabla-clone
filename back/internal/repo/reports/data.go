package reports

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

// Функция создать отзыв (БД)
func (r *Repo) RCreateReport(report *models.Report) error {
	err := r.db.QueryRow(
		`INSERT INTO reports (title, description, created_at)
		VALUES ($1, $2, CURRENT_TIMESTAMP)
		RETURNING id, created_at`,
		report.Title,
		report.Description,
	).Scan(&report.ID, &report.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

// Функция для получения отзыва по ID (БД)
func (r *Repo) RGetReportByID(id uint) (*models.Report, error) {
	var report models.Report
	err := r.db.QueryRow(
		`SELECT id, title, description, created_at
		FROM reports
		WHERE id = $1`, id,
	).Scan(&report.ID, &report.Title, &report.Description, &report.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &report, nil
}

// Функция для получения отзывов водителя (БД)
func (r *Repo) RGetReports(userID uint) ([]models.Report, error) {
	rows, err := r.db.Query(
		`SELECT id, title, description, created_at
		FROM reports WHERE user_id = $1`, userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var reports []models.Report
	for rows.Next() {
		var report models.Report
		err := rows.Scan(&report.ID, &report.Title, &report.Description, &report.CreatedAt)
		if err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}
	return reports, nil
}

// Функция для удаления отзыва (БД)
func (r *Repo) RDeleteReport(id, userID uint) error {
	_, err := r.db.Exec(
		`DELETE FROM reports
		WHERE id = $1 and user_id = $2`, id, userID,
	)
	if err != nil {
		return err
	}
	return nil
}
