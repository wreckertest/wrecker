package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
)

func CityHandler(res http.ResponseWriter, req *http.Request) {
	var str1 string = randomdata.SillyName()
	var str2 string = randomdata.Title(randomdata.Male)
	var str3 string = randomdata.FullName(randomdata.Male)
	data, _ := json.Marshal("{str1:str2,str3}")
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
