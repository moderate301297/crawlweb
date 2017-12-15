package savedata

import(
	"../dbconnection"
	"fmt"
)
// check data in mysql
func Check(title string, linkImage string) bool{
	q, err := dbconnection.Connect.Query("select title,link_image from electrics")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	var valueT string
	var valueL string
	var ok int = 0
	for q.Next() {
		err := q.Scan(&valueT,&valueL)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		if (valueT == title && valueL == linkImage) {
			ok = 1
			break;
		} else {
			continue
		}
	}
	if ok == 1 {
		return true
	} else {
		return false
	}
}