package savedata

import (
	"fmt"
	"../dbconnection"
)
// save data to mysql
func SaveData(title string, link string, linkImage string) {
	if !Check(title, linkImage) {
		_,err := dbconnection.Connect.Exec("insert electrics set title= ?, link = ?, link_image = ?", title, link, linkImage)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	} else {
		_,err := dbconnection.Connect.Exec("update electrics set title= ?, link = ?, link_image = ? where link_image = ?", title, link, linkImage, linkImage)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	} 
}
// save url error
func SaveUrlError(link string) {
	_,err := dbconnection.Connect.Exec("insert links_error set link = ?", link)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}