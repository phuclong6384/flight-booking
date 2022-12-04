package config

import (
	"github.com/spf13/viper"
)

var v = viper.New()

func Setup() {
	v.AddConfigPath(".")
	v.SetConfigName("config")
	v.SetConfigType("yml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func GetDatabaseConnection() DbConnection {
	return DbConnection{
		v.GetString("DATABASE_HOST"),
		v.GetInt("DATABASE_PORT"),
		v.GetString("DATABASE_USER"),
		v.GetString("DATABASE_PASSWORD"),
		v.GetString("DATABASE_NAME"),
	}
}

func GetUserGrpcPort() string {
	return v.GetString("USER_GRPC_PORT")
}

func GetFlightGrpcPort() string {
	return v.GetString("FLIGHT_GRPC_PORT")
}

type DbConnection struct {
	Host string
	Port int
	User string
	Pwd  string
	Db   string
}
