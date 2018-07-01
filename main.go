package main
package randomdata

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"github.com/Pallinder/go-randomdata"
)

func CityHandler(res http.ResponseWriter, req *http.Request) {
	data, _ := json.Marshal("{randomdata.SillyName():randomdata.Title(randomdata.Male),randomdata.FullName(randomdata.Male)}")
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
