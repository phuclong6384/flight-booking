package database

import "time"

const BookingStatusCreated = "Created"

type Booking struct {
	ID               int       `gorm:"column=id;not null;primaryKey"`
	CustomerUsername string    `gorm:"column=customer_username;not null"`
	FlightId         int       `gorm:"column=flight_id;not null"`
	Code             string    `gorm:"column=code;not null"`
	Status           string    `gorm:"column=status;not null"`
	ReservedSlot     int       `gorm:"reserved_slot;not null"`
	BookedDate       time.Time `gorm:"column=booked_date;not null"`
	CreatedAt        time.Time `gorm:"column=created_at;not null"`
	UpdatedAt        time.Time `gorm:"column=updated_at;not null"`
}

func (b *Booking) TableName() string {
	return "dev.booking"
}

type Flight struct {
	ID            int       `gorm:"column=id;not null;primaryKey"`
	Code          string    `gorm:"column=code;not null"`
	TotalSlot     int       `gorm:"column=total_slot;not null"`
	DepartureTime time.Time `gorm:"column=departure_time;not null"`
	ArrivalTime   time.Time `gorm:"column=arrival_time;not null"`
}

func (f *Flight) TableName() string {
	return "dev.flight"
}

type Gender int

const (
	Female Gender = iota
	Male
)

type User struct {
	Username  string    `gorm:"column=username;not null;primaryKey"`
	Password  string    `gorm:"column=password;not null"`
	FirstName string    `gorm:"column=first_name;not null"`
	LastName  string    `gorm:"column=last_name;not null"`
	Gender    Gender    `gorm:"column=gender;not null"`
	CreatedAt time.Time `gorm:"column=created_at;not null"`
	UpdatedAt time.Time `gorm:"column=updated_at;not null"`
}

func (u *User) TableName() string {
	return "dev.user"
}
