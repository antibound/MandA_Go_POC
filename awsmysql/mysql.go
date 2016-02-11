package awsmysql

import (
  _ "github.com/go-sql-driver/mysql"

  "fmt"
  "net/http"
  "database/sql"
  "log"
    "github.com/gorilla/mux"
)

type (
    // UserController represents the controller for operating on the User resource
    Mysql struct{}
)

func NewAwsController() *Mysql {
    return &Mysql{}
}

func (srv Mysql) GetRows(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)

  db, err := sql.Open("mysql", "")
  if err != nil {
    log.Println(err)
    fmt.Fprintf(w, "%s", err)
  }
  defer db.Close()

  var str string

  err = db.QueryRow("select InputOrder from Mergers").Scan(&str)
  if err != nil && err != sql.ErrNoRows {
    log.Println(err)
    fmt.Fprintf(w, "%s", err)
  }
  log.Println(str)

  fmt.Fprintf(w, "%s", str)
}

func (srv Mysql) InsertRow(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)

  db, err := sql.Open("mysql", "")
  if err != nil {
    log.Println(err)
    fmt.Fprintf(w, "%s", err)
  }
  defer db.Close()

  var str string

  vars := mux.Vars(r)
  order := vars["order"]
  sale := vars["sale"]
  margin := vars["margin"]

  _, err = db.Exec("insert into Mergers (InputOrder, Sale, Margin) Values (?, ?, ?)", order, sale, margin)
  if err != nil && err != sql.ErrNoRows {
    log.Println(err)
    fmt.Fprintf(w, "%s", err)
  }
  log.Println(str)

  fmt.Fprintf(w, "%s", str)
}
