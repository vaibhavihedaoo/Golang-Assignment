package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/aws/aws-lambda-go/lambda"
	_ "github.com/lib/pq"
	database"Golang/database"
	check"Golang/check"
)

//Structure to save data requested from database
type product_data struct {
	Id       int
	Name     string
	Category string
}
type products struct {
	Products []product_data
}
type Request struct {
	Category string
}

func Include(products_catagory []string, category string) bool {
	//check if given category belongs to specified slice of categories
	for _, v := range products_catagory {
		if v == category {
			return true
		}
	}
	return false
}

func getData(prod *products, category string) error {
	//defer db.Close()

	selectQuery := fmt.Sprintf(`select * from product where category='%s'`, category)
	rows, err := db.Query(selectQuery)

	check.checkError(err)
	defer rows.Close()
//retrive rows from the database
	for rows.Next() {
		product := product_data{}
		err := rows.Scan(&product.Id, &product.Name, &product.Category)
		check.checkError(err)
		prod.Products = append(prod.Products, product)
	}

	check.checkError(rows.Err())

	return nil
}

func displayData(data Request) (string, error) {
	//initializes database connection
	database.init()
	var result string
	values := []string{"Small", "Medium", "Large"}
	products_data := products{}

	//value := strings.Join(value, "")
	if Include(values, data.Category) {
		err = getData(&products_data, data.Category)
		check.checkError(err)

		out, err := json.Marshal(products_data)
		check.checkError(err)

		result = string(out)

	} else {
		result = "Please use proper category(Small, Medium, Large)!"
	}
	return result, nil
}

func main() {

	lambda.Start(displayData)

}
