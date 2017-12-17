package savedata

import (
	"fmt"
	"../dbconnection"
)

type DataWeb struct {
	Body string
}

// save data to mongo
func SaveDataBody(body string) {

	collection  := dbconnection.Session.DB("data_walmart").C("data")
	err := collection.Insert(&DataWeb{Body: body})
	if err != nil {
		fmt.Println(err)
    } 
}

// save data to mysql
func SaveData(title string, link string, linkImage string) {
	_,err := dbconnection.Connect.Exec("insert all_products set title= ?, link = ?, link_image = ?", title, link, linkImage)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

// save url error
func SaveUrlError(url string) {
	_,err := dbconnection.Connect.Exec("insert urls_error set url = ?", url)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
