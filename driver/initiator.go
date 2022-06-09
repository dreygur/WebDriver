package driver

import (
	"log"
	"os"
	"os/exec"
	"path"
)

func startChromeDriver() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	driverPath := path.Join(wd, "chromedriver")
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
