package main

import (
	"log"
	"net/http"
	"fmt"
)

var (
	outputFile string
)

func main() {
	http.HandleFunc("/uptime", uptimeServer)
   var err error
   var isHttps bool

   if exists("server.crt") && exists("server.key") {
      isHttps = true
   } else {
      isHttps = false
   }

   fmt.Printf("Starting server on http")
   if isHttps {
      fmt.Printf("s")
   }
   fmt.Printf("://0.0.0.0:9000\n")

   if isHttps {
	   err = http.ListenAndServeTLS("0.0.0.0:9000", "server.crt", "server.key", nil)
   } else {
	   err = http.ListenAndServe("0.0.0.0:9000", nil)
   }
	if err != nil {
		log.Fatal("ListenAndServe: " + err.Error())
	}
}
