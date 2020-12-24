package partners_db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const (
	mysql_partners_username = "mysql_partners_username"
	mysql_partners_password = "mysql_partners_password"
	mysql_partners_hostname = "mysql_partners_hostname"
	mysql_partners_schema = "mysql_partners_schema"
)

var (
	Client *sql.DB
	username = os.Getenv("mysql_partners_username")
	password = os.Getenv("mysql_partners_password")
	hostname = os.Getenv("mysql_partners_hostname")
	schema = os.Getenv("mysql_partners_schema")

)
func init() {
	datasourcename := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,password, hostname, schema,
	)

	var err error
	Client, err = sql.Open("mysql", datasourcename)

	if err != nil {
		panic(err)
	}

	if err:=Client.Ping(); err!= nil {
		panic(err)
	}

	log.Println("Database successfully connected")

}