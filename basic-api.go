package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

var db []string = []string{}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/test", test)
	r.HandleFunc("/add/{item}", addItem)
	http.ListenAndServe(":8080", r)
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		ID    string
		Name  string
		Phone string
	}{
		ID:    "01",
		Name:  "Nishant",
		Phone: "980000001",
	})
}

func addItem(w http.ResponseWriter, r *http.Request) {
	datavar := mux.Vars(r)["item"]
	db = append(db, datavar)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(db)
}
