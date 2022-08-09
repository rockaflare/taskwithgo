package database

import (
	"log"
	"taskwithgo/entity"

	"github.com/jinzhu/gorm"
)

var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Koneksi berhasil!")
	return nil
}

func Migrate(table *entity.Task) {
	Connector.AutoMigrate(&table)
	log.Println("Table dibuat.")
}
