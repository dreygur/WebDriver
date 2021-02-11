/*
 * File: methods.go
 * Created: Monday, 8th February 2021 1:48:05 am
 * Author: Rakibul Yeasin (ryeasin03@gmail.com)
 * -----
 * Last Modified: Friday, 12th February 2021 3:10:13 am
 * Modified By: Rakibul Yeasin (ryeasin03@gmail.com)
 * -----
 * Copyright (c) 2021 Slishee
 */

package webdriver

import (
	"encoding/json"
	"fmt"

	"github.com/imroc/req"
)

// Get request to an url
func Get(url string) interface{} {
	// Headers
	header := req.Header{
		"Content-Type": "application/json",
	}

	// Body
	body, err := json.Marshal(map[string]interface{}{
		"url": url,
	})
	if err != nil {
		return err.Error()
	}

	resp, err := req.Post(srv.uri+srv.sessionID+"/url", header, body)
	if err != nil {
		return err.Error()
	}
	resp.ToJSON(&res)
	// fmt.Println(resp)
	return resp
}

// Test Function
func Test() {
	fmt.Println(srv.sessionID)
}
