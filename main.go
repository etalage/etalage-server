package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := http.NewServeMux()
	log.Print("start")
	router := new(Router)
	router.BuildRoute(mux)
	server := &http.Server{
		Handler: mux,
		Addr:    "0.0.0.0:8080",
	}
	server.ListenAndServe()
	/*db := initDB()
	rows, error := db.Query("select nickname from user")
	if error != nil {
		log.Println("query error", error)
	}
	defer rows.Close()
	for rows.Next() {
		var version string
		rows.Scan(&version)
		log.Println("new line ", version)
	}*/
}

func process(response http.ResponseWriter, request *http.Request) {
	log.Print("got root")
	fmt.Fprint(response, "get your request")
}

func initDB() (db *sql.DB) {
	conn, err := sql.Open("mysql", "root:1qazZAQ!@tcp(localhost:3306)/etalage")
	if err != nil {
		log.Println("connect error")
		return
	}
	conn.SetConnMaxLifetime(3 * time.Minute)
	conn.SetMaxIdleConns(3)
	conn.SetMaxOpenConns(3)
	return conn
}
