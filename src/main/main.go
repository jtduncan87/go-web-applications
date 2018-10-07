package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	http.Handle("/", new(MyHandler))

	http.ListenAndServe(":8000", nil)
}

// MyHandler implements http Handler
type MyHandler struct {
	http.Handler
}

func (handler *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := "public/" + req.URL.Path
	fmt.Printf("Handling %s\n", path)
	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		var contentType string
		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else {
			contentType = "text/plain"
		}
		fmt.Printf("Content type is %s\n", contentType)
		w.Header().Add("Content-Type", contentType) // dash is required here for the header
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}
