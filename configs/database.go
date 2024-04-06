package configs

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "root:password@tcp(database-1.cb2okamm2pjl.us-east-1.rds.amazonaws.com:3306)/dbb?parseTime=true"

	// dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to connect to database:", err)
	}
	log.Println("Connected to database")
}

func GetDB() *gorm.DB{
	return DB
}