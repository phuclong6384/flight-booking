package dto

import (
	"time"
)

type BookFlightRequest struct {
	Username     string `json:"username" validate:"min=3"`
	Password     string `json:"password" validate:"min=3"`
	FlightId     int    `json:"flightId" validate:"required"`
	NumberOfSlot int    `json:"numberOfSlot" validate:"required"`
}

type BookingReservationResponse struct {
	ID           string    `json:"ID"`
	Username     string    `json:"username"`
	FlightId     int       `json:"flightId"`
	Code         int       `json:"code"`
	Status       string    `json:"status"`
	BookedDate   time.Time `json:"bookedDate"`
	ReservedSlot int       `json:"reserved_slot"`
}

type GetListBookingRequest struct {
	Username string `json:"username"`
}
