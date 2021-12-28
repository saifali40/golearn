package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/saifali40/golearn/cmd/pkg/config"
	"github.com/saifali40/golearn/cmd/pkg/handler"
	"github.com/saifali40/golearn/cmd/pkg/render"
)

var port = ":8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Error in creating template cache: ", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handler.NewRepository(&app)
	handler.NewHandler(repo)
	render.NewTemplates(&app)

	http.HandleFunc("/", handler.Repo.Home)
	http.HandleFunc("/about", handler.Repo.About)

	fmt.Println(fmt.Sprintf("Port number %s", port))
	_ = http.ListenAndServe(port, nil)
}
