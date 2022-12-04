package dao

import (
	"flightBooking/common/config"
	"flightBooking/common/database"

	"gorm.io/gorm"
)

type IFlightService interface {
	GetList() ([]database.Flight, error)
	Get(code string) (*database.Flight, error)
	GetById(id int) (*database.Flight, error)
}

type FlightService struct {
	DB *gorm.DB
}

func NewFlightService(conn config.DbConnection) FlightService {
	var db = database.DbConnection(conn.Host, conn.Port, conn.User, conn.Pwd, conn.Db)
	return FlightService{db}
}

func (f *FlightService) GetList() ([]database.Flight, error) {
	var query []database.Flight
	find := f.DB.Find(&query)
	if find.Error != nil {
		return nil, find.Error
	}
	return query, nil
}

func (f *FlightService) Get(code string) (*database.Flight, error) {
	query := database.Flight{}
	find := f.DB.Find(&query, database.Flight{Code: code})
	if find.Error != nil {
		return nil, find.Error
	}
	return &query, nil
}

func (f *FlightService) GetById(id int) (*database.Flight, error) {
	query := database.Flight{}
	find := f.DB.Find(&query, database.Flight{ID: id})
	if find.Error != nil {
		return nil, find.Error
	}
	return &query, nil
}
