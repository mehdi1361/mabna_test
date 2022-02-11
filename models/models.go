package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type Instrument struct {
	gorm.Model
	Name   string  `json:"name" gorm:"size:50;unique;not null"`
	Trades []Trade `json:"trades"`
}

func (i Instrument) TableName() string {
	return "instrument"
}

type Trade struct {
	gorm.Model
	DateN time.Time `json:"date_n" gorm:"size:50;unique"`
	Open  int32     `json: "open"`
	High  int32     `json: "high"`
	Low   int32     `json: "low"`
	Close int32     `json: "close"`
	InstrumentId uint    `json:"instrument_id" gorm:"Column:instrument_id"`
}

func (t Trade) TableName() string {
	return "trade"
}

func Connect() (db *gorm.DB, err error) {
	envErr := godotenv.Load()
	if envErr != nil {
		panic(envErr)
	}
	server := os.Getenv("DB_SERVER")
	database := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	conn, err := gorm.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", server, port, dbUser, database, password),
	)
	return conn, err
}

func init() {
	conn, err := Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer conn.Close()

	db := conn
	_ = db.AutoMigrate(&Instrument{}, &Trade{})
	db.Model(&Trade{}).AddForeignKey("instrument_id", "instrument(id)", "CASCADE", "CASCADE")
	if err != nil {
		fmt.Println(err)
	}

}
