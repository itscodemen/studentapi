package main

import (
	"log"
	"proj1/api/router"
	"proj1/domain/storage"
	"proj1/models"
	"proj1/pkg/db"
)

func main() {
	db, err := db.DBConfig()
	if err != nil {
		log.Fatalf("%s", err)
	}

	db.AutoMigrate(&models.Student{})

	err = storage.InitStorage(db)
	if err != nil {
		log.Fatalf("error initialize storage : %s", err)
	}

	r := router.InitRoutes()

	r.Run()
}
