package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var port = flag.String("p", "8080", "port to serve on")
var dir = flag.String("d", "/tmp", "the directory of static file to host")

// UploadFile comment
func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	f, err := os.OpenFile(*dir+"/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("File not create")
		fmt.Println(err)
		return
	}

	// Copy the file to the destination path
	io.Copy(f, file)

	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func setupRoutes() {
	http.Handle("/", http.FileServer(http.Dir(*dir)))
	http.HandleFunc("/upload", uploadFile)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func main() {
	flag.Parse()
	log.Printf("Serving %s on HTTP port: %s\n", *dir, *port)
	setupRoutes()
}
