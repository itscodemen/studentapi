package main

import (
	"flag"
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

	prt := flag.String("port", ":8080", "Port Number")
	flag.Parse()
	port := *prt
	r.Run(":" + port)
}
