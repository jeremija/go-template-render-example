package main

import (
	"html/template"
	"net/http"

	"github.com/jeremija/go-template-render-example/render"
	"github.com/jeremija/go-template-render-example/routes"
)

func Configure() http.Handler {
	render.Register(routes.IndexTemplate,
		template.Must(template.ParseFiles("templates/index.html", "templates/base.html")),
	)

	mux := http.NewServeMux()
	mux.HandleFunc(routes.IndexPath, render.Render(routes.GetIndex))
	return mux
}

func main() {
	err := http.ListenAndServe(":3000", Configure())
	if err != nil {
		panic(err)
	}
}
