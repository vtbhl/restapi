package main

import (
	"log"
	"net/http"
	"github.com/ittrada/restapi/routers"
)

func main() {
	router := routers.NewRouter()
	log.Fatal(http.ListenAndServe(":9000", router))
}
