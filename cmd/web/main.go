package main

import (
	"fmt"
	"net/http"

	"github.com/saifali40/golearn/cmd/pkg/handler"
)

var port = ":8080"

func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	fmt.Println(fmt.Sprintf("Port number %s", port))
	_ = http.ListenAndServe(port, nil)
}
