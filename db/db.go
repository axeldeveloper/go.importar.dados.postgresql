package db

import (
	"database/sql"
	"fmt"
	"log"
	//"time"
	_ "github.com/lib/pq"

)

func Conecta() *sql.DB {
	
	const (
		server   	 = "ec2-174.compute-1.amazonaws.com"
		port         = 5432
		user         = "cbogbdh"
		password     = "de0a03cd3603e165b90178d84704441"
		database     = "d3rb3aj"
		uri = "postgres://cbogbdh:de0a03cd3603e165b90178d84704441@ec2-174.compute-1.amazonaws.com:5432/d3rb3aj"
	)
	
			
	db, err := sql.Open("postgres", uri)

	if err != nil {
		log.Fatal("Open postgress connection failed:", err.Error())
		panic(err.Error())
	}

	fmt.Printf("postgress Connected!\n")

	return db
}