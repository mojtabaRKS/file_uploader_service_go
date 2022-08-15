package activities

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateConnection() *gorm.DB {
	dsn := getDsn()
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to create connection with database")
	}

	return db
}

func getDsn() string {
	db_name := os.Getenv("DATABASE_NAME")
	username := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	charset := "utf8mb4"

	return username + ":" + password + "@tcp(" + host + ":" + port + ")" + db_name + "?charset=" + charset 
}