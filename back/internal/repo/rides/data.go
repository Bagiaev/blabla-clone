package rides

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

// Функция, чтобы взять из бд поездку по id
// !!РАЗОБРАТЬСЯ С DRIVER (бд)
func (r *Repo) RGetRideByID(id int) (*models.Ride, error) {
	var ride models.Ride
	err := r.db.QueryRow(`SELECT id, created_at, driver_id, car_id, from_location, to_location, departure_at, seats_available, price_cents, description FROM rides WHERE id = $1`, id).
		Scan(&ride.ID, &ride.CreatedAt, &ride.DriverID, &ride.CarID, &ride.FromLocation, &ride.ToLocation, &ride.DepartureAt, &ride.SeatsAvailable, &ride.PriceCents, &ride.Description)

	if err != nil {
		return nil, err
	}
	return &ride, nil
}

// Функция чтобы добавить новую поездку
func (r *Repo) RCreateRide(ride *models.Ride) error {
	_, err := r.db.Exec(
		`INSERT INTO rides (created_at, driver_id, car_id, from_location, to_location, departure_at, seats_available, price_cents, description) 
		VALUES (CURRENT_TIMESTAMP, $1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at`,
		ride.DriverID,
		ride.CarID,
		ride.FromLocation,
		ride.ToLocation,
		ride.DepartureAt,
		ride.SeatsAvailable,
		ride.PriceCents,
		ride.Description,
	)
	if err != nil {
		return err
	}
	return nil
}

// Функция для получения списка поездок
func (r *Repo) RGetRides() ([]models.Ride, error) {
	rows, err := r.db.Query(
		`SELECT id, created_at, updated_at, driver_id, car_id, from_location, to_location, departure_at, seats_available, price_cents, description
		FROM rides`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var rides []models.Ride
	for rows.Next() {
		var ride models.Ride
		err := rows.Scan(&ride.ID, &ride.CreatedAt, &ride.UpdatedAt, &ride.DriverID, &ride.CarID, &ride.FromLocation, &ride.ToLocation, &ride.DepartureAt, &ride.SeatsAvailable, &ride.PriceCents, &ride.Description)
		if err != nil {
			return nil, err
		}
		rides = append(rides, ride)
	}
	return rides, nil
}

// Функция для обновления поездки по id
func (r *Repo) RUpdateRide(id int, ride *models.Ride) error {
	_, err := r.db.Exec(
		`UPDATE rides SET description = $1, updated_at = CURRENT_TIMESTAMP, price_cents = $2
		WHERE id = $3 AND driver_id = $4`,
		ride.Description, ride.PriceCents, ride.ID, ride.DriverID,
	)
	return err
}

// Функция для удаления поездки
func (r *Repo) RDeleteRide(id int) error {
	_, err := r.db.Exec(
		`DELETE FROM rides WHERE id = $1`,
		id,
	)
	return err
}
