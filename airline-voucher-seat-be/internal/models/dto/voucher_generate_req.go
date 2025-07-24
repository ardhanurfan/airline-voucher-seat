package dto

type VoucherGenerateReq struct {
	ID           string `json:"id" validate:"required"`
	Name         string `json:"name" validate:"required"`
	FlightNumber string `json:"flightNumber" validate:"required"`
	Date         string `json:"date" validate:"required"`
	Aircraft     string `json:"aircraft" validate:"required"`
}
