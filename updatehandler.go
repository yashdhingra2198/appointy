package main

import (
	"fmt"
	"database/sql"
	"strings"
	"encoding/json"
	"net/http"
	_ "github.com/lib/pq"
	
	
)

type Itemsummary struct{
	Id int    `json:"id",db:"id"`
	Value string `json:"value",db:"item_name"`
	Completed bool  `json:"completed",db:"status"`
}


func updateitemHandler(w http.ResponseWriter, r *http.Request){
	// Parse and decode the request body into a new `Credentials` instance
	item := &Itemsummary{}

	err := json.NewDecoder(r.Body).Decode(item)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return 
	}
	
sqlStatement1 := `
SELECT item_name,list_name
FROM items
WHERE id = $1;`
var name string
var list_name string
row := db.QueryRow(sqlStatement1, item.Id)
err2 := row.Scan(&name,&list_name)
switch err2 {
case sql.ErrNoRows:
  fmt.Println("No rows were returned!")
  return
case nil:
  fmt.Println(user)
default:
  panic(err2)
}


sqlStatement2 := `
UPDATE items
SET item_name = $2,status=$3
WHERE id = $1;`
row= db.QueryRow(sqlStatement2,item.Id ,item.Value,item.Completed)

fmt.Println(name)

if name!=item.Value{

sqlStatement3 := `
SELECT list_items
FROM list
WHERE list_name = $1;`
var line string
row = db.QueryRow(sqlStatement3, list_name)
err2 = row.Scan(&line)
switch err2 {
case sql.ErrNoRows:
  fmt.Println("No rows were returned!")
  return
case nil:
  fmt.Println(user)
default:
  panic(err2)
}
fmt.Println(name)
fmt.Println(line)






	// Split the line on commas.
        parts := strings.Split(line, ",")

        // Loop over the parts from the string.
        for i := range parts {
            if parts[i]==name{
            	parts[i]=item.Value
            }
        }

        var out string

        for i := range parts {
            out=out+parts[i]+","
        }

fmt.Println(out)


        sqlStatement4 := `
		UPDATE list
		SET list_items = $2
		WHERE list_name = $1;`
		row= db.QueryRow(sqlStatement4,list_name ,parts)


}



}
