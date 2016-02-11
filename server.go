package main

import (
    // Standard library packages
    "net/http"
    "log"

    // Third party packages
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"

    // Custom packages
    "manda/controllers"
    "manda/awsmysql"
    "html/template"
)

func main() {
      // Get a UserController instance
      uc := controllers.NewUserController()
      srv := awsmysql.NewAwsController()

      r := mux.NewRouter().StrictSlash(true)
      r.HandleFunc("/user", uc.GetUser).Methods("GET")
      r.HandleFunc("/insert/{order}/{sale}/{margin}", srv.InsertRow).Methods("POST")
      r.HandleFunc("/rows", srv.GetRows).Methods("GET")
    	r.HandleFunc("/", index).Methods("GET")

    	http.Handle("/", r)

    	log.Println("Listening at 8080")
    	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Index request", r.RemoteAddr, r.URL)
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}
