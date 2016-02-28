package main

import (
	"encoding/json"
	"net/http"
)

type book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	http.HandleFunc("/", showBooks)
	http.ListenAndServe(":8080", nil)
}

func showBooks(w http.ResponseWriter, r *http.Request) {
	b := book{"Building Web Apps with Go", "Jeremy Saenz"}

  w.Header().Set("Content-Type", "application/json")

  enc := json.NewEncoder(w)
  if err := enc.Encode(b); err != nil { //combining these two lines ensures err is only within this conditional's scope. indicates intention. err not needed below.
    fmt.println("Error encoding!", err)
  }
	// js, err := json.Marshal(b)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.Write(js)
}
