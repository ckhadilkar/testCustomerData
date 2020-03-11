package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "us-ss-dbaadm-6281-snapshot-userdb.enova.com"
	port     = 5432
	user     = "cnuapp"
	password = "gouwuiv2"
	dbname   = "cnuapp_prod"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	paydayPaidOffCustomer(*db)
	locIssuedCustomer(*db)
}