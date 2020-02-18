package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/esviderskii/go-server/helpers"
	"github.com/spf13/viper"
)

var port = flag.String("p", "8080", "port to serve on")
var dir = flag.String("d", "./", "the dir of static file to host")

func setupRoutes() {

	http.Handle("/", http.FileServer(http.Dir(viper.GetString("path"))))
	http.HandleFunc("/upload", helpers.UploadFile)
	log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), nil))
}

func main() {
	flag.Parse()
	log.Printf("Serving %s on HTTP port: %s\n", viper.GetString("path"), viper.GetString("port"))
	setupRoutes()
}
