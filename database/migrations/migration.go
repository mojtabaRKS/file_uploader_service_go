package migrations

import (
	database "file-uploader/internal/activities/database"
	"log"
)

func Migrate() {
	db := database.CreateConnection()
	
	err := db.AutoMigrate(&File{})

	if err != nil {
		log.Println(err)
		panic("cant migrate")
	}
}