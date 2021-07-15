package main

import (
	"database/sql"
	"fmt"
		_"github.com/lib/pq"
	"github.com/aws/aws-lambda-go/lambda"
	database "Golang/database"
  mail "Goalng/mail"
	check"Golang/check"
)
//Structure to save data from request body
type MyRequest struct{
	id int
	name string
  category string
}

type MyResponse struct{
	message string `json:"Inserted:"`
}

//function to insert product into the table
func insertData(event MyRequest)(MyResponse,error)  {
	//initializes database connection
database.init()
//query to insert product
  selectQuery := fmt.Sprintf(`INSERT INTO product(id, name,category)
VALUES ($1, $2, $3)
RETURNING id`)
   rows,err := p.db.Query(selectQuery,event.id,event.name,event.category)
check.checkError(rows.Err())
  //defer rows.Close()
	mail.sendMail()
return  MyResponse{message: fmt.Sprintf("%s is inserted and mail sent", event.id)}, nil

}

func main() {
lambda.Start(insertData)

}
