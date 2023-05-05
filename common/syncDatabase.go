package common

import (
	"authen-author-example/models"
	"log"
)

func SyncDatabase() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed!")
}
