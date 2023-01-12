package helpers

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB_USER = "postgres"
var DB_PASSWORD = "Paavani"
var DB_NAME = "pro_cat_shop_cart_service_test"

func SetupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		CheckErr(err)
	}

	return db
}
