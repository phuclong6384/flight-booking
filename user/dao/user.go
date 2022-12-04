package dao

import (
	"flightBooking/common/config"
	"flightBooking/common/database"
	"flightBooking/common/util"

	"gorm.io/gorm"
)

type IUserService interface {
	Create(user *database.User) (*database.User, error)
	Update(user *database.User) (*database.User, error)
	GetByUsername(username string) (*database.User, error)
	ValidatePassword(user *database.User, password string) bool
}

type UserService struct {
	IUserService
	DB *gorm.DB
}

func NewUserService(conn config.DbConnection) UserService {
	var db = database.DbConnection(conn.Host, conn.Port, conn.User, conn.Pwd, conn.Db)
	return UserService{DB: db}
}

func (u *UserService) Create(user *database.User) (*database.User, error) {
	create := u.DB.Create(user)
	if create.Error != nil {
		return user, create.Error
	}
	return user, nil
}

func (u *UserService) Update(user *database.User) (*database.User, error) {
	save := u.DB.Save(user)
	if save.Error != nil {
		return user, save.Error
	}
	return user, nil
}

func (u *UserService) GetByUsername(username string) (*database.User, error) {
	query := database.User{}
	find := u.DB.Find(&query, database.User{Username: username})
	if find.Error != nil {
		return nil, find.Error
	}
	return &query, nil
}

func (u *UserService) ValidatePassword(user *database.User, password string) bool {
	return util.ValidatePassword(password, user.Password)
}
