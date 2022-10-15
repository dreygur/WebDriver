package driver

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
)

func startChromeDriver() {
	var driverName string = "chromedriver"
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	osName := runtime.GOOS
	if osName == "windows" {
		driverName = fmt.Sprintf("%s.exe", driverName)
	}
	driverPath := path.Join(wd, driverName)
	if _, err := os.Stat(driverPath); err != nil {
		err = download()
		if err != nil {
			panic(err)
		}
	}

	// Start ChromeDriver
	go func() {
		cmd := exec.Command(driverPath, "--port=3535")
		if err := cmd.Start(); err != nil {
			return
		}
		if err := cmd.Wait(); err != nil {
			return
		}
		defer cmd.Process.Kill()
	}()
}
