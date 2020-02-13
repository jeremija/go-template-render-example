package render

import (
	"context"
	"net/http"
)

type Meta struct {
	Title       string
	ImageURL    string
	URL         string
	Description string
}

type MetaContext map[string]interface{}

const (
	MetaContextKey = "MetaContext"

	MetaKey = "Meta"
)

func WithMetaContext(r *http.Request) *http.Request {
	c := MetaContext{
		MetaKey: Meta{},
	}
	ctx := context.WithValue(r.Context(), MetaContextKey, c)
	return r.WithContext(ctx)
}

func getParams(r *http.Request) MetaContext {
	params, ok := r.Context().Value(MetaContextKey).(MetaContext)
	if !ok {
		panic("No render params set in request context")
	}
	return params
}

func SetParam(r *http.Request, key string, value interface{}) {
	getParams(r)[key] = value
}

func GetParam(r *http.Request, key string) (interface{}, bool) {
	item, ok := getParams(r)[key]
	return item, ok
}

func SetMeta(r *http.Request, meta Meta) {
	SetParam(r, MetaKey, meta)
}

func GetMeta(r *http.Request) (meta Meta) {
	value, ok := GetParam(r, MetaKey)
	if !ok {
		return
	}
	v, ok := value.(Meta)
	if !ok {
		return
	}
	meta = v
	return
}
