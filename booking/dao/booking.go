package dao

import (
	"flightBooking/common/config"
	"flightBooking/common/database"
	"log"

	"gorm.io/gorm"
)

type IBookingService interface {
	GetList(username string) ([]database.Booking, error)
	GetById(id int) (*database.Booking, error)
	Create(booking *database.Booking) (*database.Booking, error)
	GetByFlightId(flightId int) ([]database.Booking, error)
}

type BookingService struct {
	DB *gorm.DB
}

func NewBookingService(conn config.DbConnection) BookingService {
	var db = database.DbConnection(conn.Host, conn.Port, conn.User, conn.Pwd, conn.Db)
	return BookingService{db}
}

func (b *BookingService) GetList(username string) ([]database.Booking, error) {
	var booking []database.Booking
	find := b.DB.Find(&booking, database.Booking{CustomerUsername: username})
	if find.Error != nil {
		return nil, find.Error
	}
	return booking, nil
}

func (b *BookingService) GetById(id int) (*database.Booking, error) {
	booking := database.Booking{}
	find := b.DB.Find(&booking, database.Booking{ID: id})
	if find.Error != nil {
		return nil, find.Error
	}
	return &booking, nil
}

func (b *BookingService) Create(booking *database.Booking) (*database.Booking, error) {
	create := b.DB.Create(booking)
	if create.Error != nil {
		return nil, create.Error
	}
	return booking, nil
}

func (b *BookingService) GetByFlightId(flightId int) ([]database.Booking, error) {
	var result []database.Booking
	find := b.DB.Find(&result, database.Booking{FlightId: flightId})
	if find.Error != nil {
		log.Println("Error occurred in GetByFlightId", find.Error)
		return nil, find.Error
	}
	return result, nil
}
