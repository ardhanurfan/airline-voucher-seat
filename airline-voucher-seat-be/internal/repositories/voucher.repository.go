package repositories

import (
	"airline-voucher-seat-be/internal/models/domain"
	"airline-voucher-seat-be/internal/models/dto"
	"context"
	"database/sql"
	"errors"
)

type VoucherRepositoryInterface interface {
	Create(ctx context.Context, voucher domain.Voucher) error
	CheckVoucherIsExist(ctx context.Context, checkReq dto.VoucherCheckReq) (bool, error)
	GetVouchers(ctx context.Context) ([]domain.Voucher, error)
}

type VoucherRepository struct {
	DB *sql.DB
}

func NewVoucherRepository(db *sql.DB) VoucherRepositoryInterface {
	return &VoucherRepository{DB: db}
}

func (r *VoucherRepository) Create(ctx context.Context, voucher domain.Voucher) error {
	stmt := `INSERT INTO vouchers (
		crew_name, crew_id, flight_number, flight_date,
		aircraft_type, seat1, seat2, seat3, created_at
	  ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.DB.Exec(stmt,
		voucher.CrewName,
		voucher.CrewID,
		voucher.FlightNumber,
		voucher.FlightDate,
		voucher.AircraftType,
		voucher.Seat1,
		voucher.Seat2,
		voucher.Seat3,
		voucher.CreatedAt,
	)
	if err != nil {
		return errors.New("[VoucherRepository] failed to create voucher: " + err.Error())
	}

	return nil
}

func (r *VoucherRepository) CheckVoucherIsExist(ctx context.Context, checkReq dto.VoucherCheckReq) (bool, error) {
	stmt := `SELECT COUNT(*) FROM vouchers WHERE flight_number = ? AND flight_date = ?`
	var count int
	err := r.DB.QueryRowContext(ctx, stmt, checkReq.FlightNumber, checkReq.Date).Scan(&count)
	if err != nil {
		return false, errors.New("[VoucherRepository] failed to check voucher existence: " + err.Error())
	}

	return count > 0, nil
}

func (r *VoucherRepository) GetVouchers(ctx context.Context) ([]domain.Voucher, error) {
	stmt := `SELECT * FROM vouchers`
	rows, err := r.DB.QueryContext(ctx, stmt)
	if err != nil {
		return nil, errors.New("[VoucherRepository] failed to get vouchers: " + err.Error())
	}
	defer rows.Close()

	vouchers := []domain.Voucher{}
	for rows.Next() {
		var voucher domain.Voucher
		if err := rows.Scan(
			&voucher.ID,
			&voucher.CrewName,
			&voucher.CrewID,
			&voucher.FlightNumber,
			&voucher.FlightDate,
			&voucher.AircraftType,
			&voucher.Seat1,
			&voucher.Seat2,
			&voucher.Seat3,
			&voucher.CreatedAt,
		); err != nil {
			return nil, errors.New("[VoucherRepository] failed to scan voucher: " + err.Error())
		}

		vouchers = append(vouchers, voucher)
	}

	return vouchers, nil
}
