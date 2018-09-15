package main

import (
	"database/sql"
	"net/http"
	"log"
	_ "github.com/lib/pq"
	"fmt"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "9205607899"
  dbname   = "appointy"
)


var db *sql.DB

func main() {
	
	http.HandleFunc("/todolist:updateItem", updateitemHandler)
	
	// initialize our database connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
    var err error
      db, err = sql.Open("postgres", psqlInfo)
      if err != nil {
        panic(err)
      }
      defer db.Close()

      err = db.Ping()
      if err != nil {
        panic(err)
      }
	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":9090", nil))
}

