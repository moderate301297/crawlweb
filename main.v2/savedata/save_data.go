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
	}
}
// save link
func SaveLink(link string) {
	_,err := dbconnection.Connect.Exec("insert links set link = ?", link)
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
// save url product error
func SaveLinkError(url string) {
	_,err := dbconnection.Connect.Exec("insert links_error set url = ?", url)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}