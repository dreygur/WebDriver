package driver

import (
	"fmt"
)

// Takes a Screenshot of the current page
func (d *WebDriver) Screenshot() (string, error) {
	r, err := get(fmt.Sprintf("%s/%s/screenshot", d.uri, d.session))
	if err != nil {
		return "", err
	}
	return r, nil
}
