package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
)

func CityHandler(res http.ResponseWriter, req *http.Request) {
	data, _ := json.Marshal("{'cities':'San Francisco, Amsterdam, Berlin, New York','Tokyo'}")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}

func main() {
	fmt.Println("hello world")
	http.HandleFunc("/cities.json", CityHandler)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
