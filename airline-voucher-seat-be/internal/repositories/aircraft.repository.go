package repositories

import (
	"airline-voucher-seat-be/internal/models/domain"
	"context"
	"database/sql"
	"errors"
)

type AircraftRepositoryInterface interface {
	GetSeatLayoutByAircraft(ctx context.Context, aircraftType string) (domain.SeatLayout, error)
	GetAircraftTypes(ctx context.Context) ([]string, error)
}

type AircraftRepository struct {
	DB *sql.DB
}

func NewAircraftRepository(db *sql.DB) AircraftRepositoryInterface {
	return &AircraftRepository{DB: db}
}

func (r *AircraftRepository) GetSeatLayoutByAircraft(ctx context.Context, aircraftType string) (domain.SeatLayout, error) {
	stmt := `SELECT aircraft_type, row_start, row_end, seats_per_row FROM aircrafts WHERE aircraft_type = ?`
	var layout domain.SeatLayout

	err := r.DB.QueryRowContext(ctx, stmt, aircraftType).Scan(&layout.AircraftType, &layout.RowStart, &layout.RowEnd, &layout.SeatsPerRow)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.SeatLayout{}, errors.New("[AircraftRepository] no seat layout found for aircraft type: " + aircraftType)
		}
		return domain.SeatLayout{}, errors.New("[AircraftRepository] failed to get seat layout: " + err.Error())
	}

	return layout, nil
}

func (r *AircraftRepository) GetAircraftTypes(ctx context.Context) ([]string, error) {
	stmt := `SELECT DISTINCT aircraft_type FROM aircrafts`
	var aircraftTypes []string

	rows, err := r.DB.QueryContext(ctx, stmt)
	if err != nil {
		return nil, errors.New("[AircraftRepository] failed to get aircraft types: " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var aircraftType string
		if err := rows.Scan(&aircraftType); err != nil {
			return nil, errors.New("[AircraftRepository] failed to scan aircraft type: " + err.Error())
		}
		aircraftTypes = append(aircraftTypes, aircraftType)
	}
	return aircraftTypes, nil
}
