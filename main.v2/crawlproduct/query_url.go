package crawlproduct

import (
	"fmt"
	"../dbconnection"
)

func QueryLink() {
	q, err := dbconnection.Connect.Query("select url from links")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	var value string
	for q.Next() {
		err := q.Scan(&value)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		CrawlProduct(value)
	}
}