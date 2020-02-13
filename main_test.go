package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jeremija/go-template-render-example/routes"
)

func TestGetIndex(t *testing.T) {
	h := Configure()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", routes.IndexPath, nil)

	h.ServeHTTP(w, r)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("Unexpected status code: %d", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %s", err)
	}
	expected := `<!DOCTYPE html>
<html>
<head>
  <meta property="og:title" content="My Title">
  <meta property="og:description" content="My Description">
  
  <meta property="og:url" content="/index">
  <meta name="twitter:title" content="My Title">
  <meta name="twitter:description" content="My Description">
</head>
<body>

<h1>My Title</h1>
<p>5</p>

</body>
</html>`
	result := strings.Trim(string(b), "\n ")
	if result != expected {
		t.Fatalf("Expected response to be:\n===\n%s===\nbut was:\n===\n%s===", expected, result)
	}
}
