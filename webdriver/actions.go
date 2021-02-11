/*
 * File: actions.go
 * Created: Friday, 12th February 2021 12:29:32 am
 * Author: Rakibul Yeasin (ryeasin03@gmail.com)
 * -----
 * Last Modified: Friday, 12th February 2021 3:10:13 am
 * Modified By: Rakibul Yeasin (ryeasin03@gmail.com)
 * -----
 * Copyright (c) 2021 Slishee
 */

package webdriver

import (
	"github.com/imroc/req"
)

// Screenshot of a webpage
func Screenshot(name string) interface{} {
	// Headers
	header := req.Header{
		"Content-Type": "application/json",
	}

	resp, err := req.Get(srv.uri+srv.sessionID+"/screenshot/", header)
	if err != nil {
		return err.Error()
	}
	resp.ToJSON(&res)

	// reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(res["value"].(string)))
	// m, _, err := image.Decode(reader)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// //Encode from image format to writer
	// pngFilename := name + ".png"
	// f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return err.Error()
	// }

	// err = png.Encode(f, m)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return err.Error()
	// }
	// fmt.Println("Png file", pngFilename, "created")

	return resp
}
