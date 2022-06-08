package driver

import (
	"fmt"
	"net/http"
)

func (d *WebDriver) Get(url string) (*http.Response, error) {
	uri := fmt.Sprintf("%s/%s/url", d.uri, d.session)
	r, err := post(uri, []byte(`{"url": "`+url+`"}`))
	if err != nil {
		println(err)
		return nil, err
	}
	return r, nil
}
