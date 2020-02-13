package routes

import (
	"net/http"

	"github.com/jeremija/go-template-render-example/render"
)

const IndexTemplate = "index.html"
const IndexPath = "/index"

type IndexData struct {
	Value int
}

func GetIndex(w http.ResponseWriter, r *http.Request) (string, interface{}, error) {
	render.SetMeta(r, render.Meta{
		Title:       "My Title",
		Description: "My Description",
	})
	return IndexTemplate, IndexData{Value: 5}, nil
}
