package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "strings"
    "encoding/json"
    "strconv"

    _ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "9205607899"
  dbname   = "appointy"
)


type errRepoNotInitialized string

func (e errRepoNotInitialized) Error() string {
    return string(e)
}

type errRepoNotFound string

func (e errRepoNotFound) Error() string {
    return string(e)
}


type itemSummary struct {
    ID         int
    Name       string
    Status     bool

}
type items struct {
    Items []itemSummary
}
var db *sql.DB
func main() {
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

      
        
    
    
    http.HandleFunc("/todolist:getItem/",getitemHandler)
   
    log.Fatal(http.ListenAndServe(":9090", nil))


    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }


}




func parseParams(req *http.Request, prefix string, num int) ([]string, error) {
    url := strings.TrimPrefix(req.URL.Path, prefix)
    params := strings.Split(url, "/")
    if len(params) != num || len(params[0]) == 0  {
        return nil, fmt.Errorf("Bad format. Expecting exactly %d params", num)
    }
    return params, nil
}


func getitemHandler(w http.ResponseWriter, req *http.Request) {
    
    item := itemSummary{}
    params, err := parseParams(req, "/todolist:getItem/", 1)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }
    
    

     i1, err := strconv.Atoi(params[0])
    if err == nil {
        fmt.Println(i1)
    }
    item.ID = i1


    data, err := queryItem(&item)
    if err != nil {
        switch err.(type) {
        case errRepoNotFound:
            http.Error(w, err.Error(), 404)
        case errRepoNotInitialized:
            http.Error(w, err.Error(), 401)
        default:
            http.Error(w, err.Error(), 500)
        }
        return
    }

    out, err := json.Marshal(data)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    fmt.Fprintf(w, string(out))
}


func queryItem(item *itemSummary) (*itemSummary, error) {
    data,err := fetchItem(item)
    if err != nil {
        return nil, err
    }

    return data,nil
}



func fetchItem(item *itemSummary) (*itemSummary,error) {
    
    
    sqlStatement := `
        SELECT
            id,item_name,status
        FROM items
        WHERE id=$1
        LIMIT 1;`
    row := db.QueryRow(sqlStatement,item.ID)
    err := row.Scan(&item.ID, &item.Name, &item.Status,)
    if err != nil {
        switch err {
        case sql.ErrNoRows:
          
            return nil,errRepoNotFound("Item not found")
        default:
            return nil,err
        }
    }
    
    return item,nil
}
