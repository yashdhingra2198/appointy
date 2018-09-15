package main

import (
	
	"database/sql"
	"encoding/json"
	"net/http"
	_ "github.com/lib/pq"
	"fmt"
	
)

type Itemsummary struct{
	Id int    `json:"id",db:"id"`
	Value string `json:"value",db:"item_name"`
	Completed bool  `json:"completed",db:"status"`
}


type Item struct {

	List_id int `json:"list_id"`
	Itemdetails  Itemsummary `json:"item"`
	
}
type List struct{
	List_id int
	List_name string
	List_items string

}

func additemHandler(w http.ResponseWriter, r *http.Request){
	// Parse and decode the request body into a new `Credentials` instance
	item := &Item{}

	err := json.NewDecoder(r.Body).Decode(item)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return 
	}
	sqlStatement := `SELECT * FROM list WHERE list_id=$1;`
	var list List
	row := db.QueryRow(sqlStatement, item.List_id)
	err2 := row.Scan(&list.List_id, &list.List_name, &list.List_items)
	switch err2 {
	case sql.ErrNoRows:
	  fmt.Println("No rows were returned!")
	  return
	case nil:
	  fmt.Println(user)
	default:
	  panic(err2)
	}
	list.List_items=list.List_items+","+item.Itemdetails.Value
	sqlStatement2 := `
	UPDATE list
	SET list_items = $2
	WHERE list_id = $1;`
	_, err = db.Exec(sqlStatement2,item.List_id ,list.List_items)
	if err != nil {
	  panic(err)
	}
	
	if _, err = db.Query("insert into items values ($1, $2, $3,$4)", item.Itemdetails.Id,item.Itemdetails.Value,item.Itemdetails.Completed ,list.List_name); err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}




	


}
