package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//Index is ...
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(map[string]string{
		"addr": r.RemoteAddr,
	})

	w.Write(body)
}

//ReverseProxy is ...
func ReverseProxy(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:8080"
	if "" != os.Getenv("URL") {
		url = os.Getenv("URL")
	}
	log.Printf("Target %s.", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	println("You are " + string(body))
	w.Write(body)
}
