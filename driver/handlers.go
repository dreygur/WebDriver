package driver

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Generate session ID and save it to Webdriver session
func getSession() string {
	c := http.Client{}
	data, err := json.Marshal(map[string]map[string]map[string]bool{
		"capabilities": {
			"alwaysMatch": {
				"acceptInsecureCerts": true,
			},
		},
	})
	if err != nil {
		panic(err)
	}
	r, err := c.Post(uri, "application/json", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	var session map[string]map[string]string
	json.NewDecoder(r.Body).Decode(&session)
	return session["value"]["sessionId"]
}

// Send a POST request to Webdriver server
func post(url string, data []byte) (*http.Response, error) {
	c := http.Client{}
	r, err := c.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return r, nil
}

// Send a GET request to Webdriver server
func get(url string) (string, error) {
	c := http.Client{}
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-Type", "application/json")
	res, err := c.Do(r)
	if err != nil {
		print(err)
		return "", err
	}
	defer res.Body.Close()
	var data map[string]string
	json.NewDecoder(res.Body).Decode(&data)
	return string(data["value"]), nil
}
