package configs

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "host=dcc.cb2okamm2pjl.us-east-1.rds.amazonaws.com user=postgres password=Vincent27 dbname=DCC port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	// dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to connect to database")
	}
	log.Println("Connected to database")
}

func GetDB() *gorm.DB{
	return DB
}