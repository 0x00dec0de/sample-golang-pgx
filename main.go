package main

import (
	"log"

	"github.com/jackc/pgx"
)

func main() {
	log.Println("start sample")

	pgxConfig := pgx.ConnConfig{
		Host:     "localhost",
		Database: "quetest",
		User:     "quetest",
	}
	pgxConnPoolConfig := pgx.ConnPoolConfig{pgxConfig, 3, nil}
	conn, err := pgx.NewConnPool(pgxConnPoolConfig)
	if err != nil {
		log.Fatal(err)
	}

	var n int32
	err = conn.QueryRow("select id from item limit 1").Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("payment_id: %d", n)
}
