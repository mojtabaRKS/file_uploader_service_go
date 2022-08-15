package file_uploader

import (
	"file-uploader/database/migrations"
	database "file-uploader/internal/activities/database"
)

func Save(file_name string, size int) (migrations.File, error) {
	db := database.CreateConnection()

	model := migrations.File{
		Path:     "storage/files/" + file_name,
		FileName: file_name,
		Size:     uint(size),
	}

	result := db.Create(&model)

	if result.Error != nil {
		return migrations.File{}, result.Error
	}

	return model, nil
}
