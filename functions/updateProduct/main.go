package main

import (
  "database/sql"
  "fmt"
  "github.com/aws/aws-lambda-go/lambda"
  _"github.com/lib/pq"
  database"Golang/database"
  check"Golang/check"
)

//Structure to save data requested
type MyEvent struct {
        Id int     `json:"Enter id"`
        Name string `json:"Enter name of product"`
        Category string `json:"Enter updated category of product"`
}

type MyResponse struct {
        Message string `json:"Updated Product:"`
}

//function to update category of product
func updateData(event MyEvent)(MyResponse,error)  {
//initializes database connection
database.init()
//query to update a product
	query :=
	 `UPDATE product
	 SET category = ($1)
	   where id = ($2) and name = ($3)`
stmt,err := database.db.Prepare(query)
result,err := stmt.Exec(event.Category,event.Id,event.Name)
  check.checkError(err)
  //defer rows.Close()
  return MyResponse{Message: fmt.Sprintf("product at id %d is updated: %s !", event.Id, event.Name,result)}, nil
}

func main() {
        lambda.Start(updateData)
}
