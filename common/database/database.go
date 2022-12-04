package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func DbConnection(host string, port int, user string, password string, dbName string) *gorm.DB {
	connectionString := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		host, user, password, dbName, port)
	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  connectionString,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "dev.",
			SingularTable: false,
		}})
	if err != nil {
		panic("Cannot establish connection to database")
	}
	return db
}
