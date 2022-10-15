package main

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/dreygur/webdriver/driver"
)

func main() {
	d := driver.InitWebdriver()

	_, err := d.Get("http://www.youtube.com")
	if err != nil {
		log.Fatal(err)
		return
	}

	s, err := d.Screenshot()
	if err != nil {
		log.Fatal(err)
		return
	}

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(s))
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	//Encode from image format to writer
	pngFilename := "screenshot.png"
	f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = png.Encode(f, m)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Png file", pngFilename, "created")
}
