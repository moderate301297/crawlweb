package dbconnection

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2"
)

var (
	Connect *sql.DB
	Session *mgo.Session
)

func init() {
	var err error
	Connect, err = sql.Open("mysql", "root:30121996@tcp(localhost:3306)/data_walmart")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	Session, err = mgo.Dial("localhost:27017")
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	Session.SetMode(mgo.Monotonic, true)
}

func Close() {
	Connect.Close()
	Session.Close()
}