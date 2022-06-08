package driver

import "net/http"

type WebDriver struct {
	client  *http.Client
	session string
	uri     string
}

type Header struct {
	Name string
}
