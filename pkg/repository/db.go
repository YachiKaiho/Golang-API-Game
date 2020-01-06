package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// blank import for MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

// DriverName
const driverName = "mysql"

// DB use in each repository
var DB *sql.DB

func init() {
	/* ===== connect database ===== */
	// user
	user := os.Getenv("MYSQL_USER")
	// password
	password := os.Getenv("MYSQL_PASSWORD")
	// connect host
	host := os.Getenv("MYSQL_HOST")
	// connect port
	port := os.Getenv("MYSQL_PORT")
	// database
	database := os.Getenv("MYSQL_DATABASE")

	// connect info
	// user:password@tcp(host:port)/database
	var err error
	DB, err = sql.Open(driverName,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database))
	if err != nil {
		log.Fatal(err)
	}
}
