package dto

type VoucherCheckReq struct {
	FlightNumber string `json:"flightNumber" validate:"required"`
	Date         string `json:"date" validate:"required"`
}
