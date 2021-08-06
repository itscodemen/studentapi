package main

import (
	"flag"
	"log"
	"proj1/api/router"
	"proj1/domain/storage"
	"proj1/pkg/db"
)

func main() {
	port := flag.String("port", "8080", "Port Number")
	flag.Parse()

	db, err := db.DBConfig()
	if err != nil {
		log.Fatalf("%s", err)
	}

	err = storage.InitStorage(db)
	if err != nil {
		log.Fatalf("error initialize storage : %s", err)
	}

	r := router.InitRoutes()

	r.Run(":" + *port)
}
