package database

import (
	"database/sql"
	"fmt"
	_ "github.com/aws/aws-lambda-go/lambda"
	_ "github.com/lib/pq"
)

const (
	host     = "product-database.cncz6v0mbvgf.ap-south-1.rds.amazonaws.com"
	port     = 5432
	user     = "product_user"
	password = "product_1"
	dbname   = "product"
)

var db *sql.DB
var err error

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	//initializes database connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	checkError(err)
	err = db.Ping()
	checkError(err)
	fmt.Println("\nSuccessfully connected!")

}
