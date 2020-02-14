package render

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

type PageHandler func(w http.ResponseWriter, r *http.Request) (templateName string, data interface{}, err error)

const (
	ErrRenderingPage = "Error rendering page"
)

var templates = map[string]*template.Template{}

func Register(templateName string, template *template.Template) {
	templates[templateName] = template
}

func Render(h PageHandler) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r = WithMetaContext(r)
		templateName, data, err := h(w, r)
		if err != nil {
			fmt.Printf("Error handling request: %s\n", err)
			http.Error(w, ErrRenderingPage, http.StatusInternalServerError)
			return
		}

		template, ok := templates[templateName]
		if !ok {
			fmt.Printf("No template not found: %s\n", templateName)
			http.Error(w, ErrRenderingPage, http.StatusInternalServerError)
			return
		}

		meta := GetMeta(r)
		if meta.URL == "" {
			meta.URL = r.URL.EscapedPath()
		}

		dataMap := map[string]interface{}{
			"Data": data,
			"Meta": meta,
		}

		// TODO write to buffer pool instead of directly to http
		var b bytes.Buffer
		err = template.Execute(&b, dataMap)
		if err != nil {
			fmt.Printf("Error executing template: %s\n", err)
			http.Error(w, ErrRenderingPage, http.StatusInternalServerError)
			return
		}
		io.Copy(w, &b)
	}
	return http.HandlerFunc(fn)
}
