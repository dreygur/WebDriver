package driver

import (
	"net/http"
	"time"
)

var (
	// Webdriver Server URI
	uri string = "http://localhost:3535/session"
)

// Initialize Webdriver session
func InitWebdriver() Driver {
	startChromeDriver()
	time.Sleep(time.Second * 2)

	return &WebDriver{
		client:  &http.Client{},
		session: getSession(),
		uri:     uri,
	}
}

// Webdriver session interface
type Driver interface {
	// Method to browse a webpage
	Get(url string) (*http.Response, error)
	// Takes a Screenshot of the current page
	Screenshot() (string, error)
}
