package driver

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

// Download ChromeDriver
func download() error {
	version, dLoadableZip := getVersion()
	// zipPath := path.Join(os.Getenv("HOME"), ".webdriver", "chromedriver.zip")
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	zipPath := path.Join(wd, "chromedriver.zip")
	driverPath := path.Join(wd, "chromedriver")
	if _, err := os.Stat(driverPath); err != nil {
		err = downloadChromeDriver(version, dLoadableZip, zipPath)
		if err != nil {
			return err
		}
		os.Remove(zipPath)
	}

	return nil
}

func downloadChromeDriver(version, dLoadableZip, zipPath string) error {
	fmt.Println("Downloading ChromeDriver...")
	dLoadUri := fmt.Sprintf("https://chromedriver.storage.googleapis.com/%s/chromedriver_%s.zip", version, dLoadableZip)

	resp, err := http.Get(dLoadUri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	// Unzip
	err = unzipSource(zipPath, ".")
	if err != nil {
		return err
	}

	return nil
}

func getVersion() (string, string) {
	// Get Operating System
	osName := runtime.GOOS
	var (
		chromeVersion string
		dLoadableZip  string
		err           error
		res           []byte
		r             = regexp.MustCompile(`(\d+\.\d+\.\d+\.\d+)`)
	)

	switch osName {
	case "windows":
		out := exec.Command("powershell", `(Get-Item (Get-ItemProperty 'HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\App Paths\chrome.exe').'(Default)').VersionInfo`)
		res, err = out.Output()
		if err != nil {
			log.Fatal(err)
		}
		chromeVersion = r.FindString(string(res))
		dLoadableZip = "win32"
	case "linux":
		out := exec.Command("google-chrome-stable", "--version")
		res, err = out.Output()
		if err != nil {
			log.Fatal(err)
		}
		chromeVersion = r.FindString(string(res))
		dLoadableZip = "linux64"
	}

	return chromeVersion, dLoadableZip
}

func unzipSource(source, destination string) error {
	// 1. Open the zip file
	reader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	// 2. Get the absolute destination path
	destination, err = filepath.Abs(destination)
	if err != nil {
		return err
	}

	// 3. Iterate over zip files inside the archive and unzip each of them
	for _, f := range reader.File {
		err := unzipFile(f, destination)
		if err != nil {
			return err
		}
	}

	return nil
}

func unzipFile(f *zip.File, destination string) error {
	// 4. Check if file paths are not vulnerable to Zip Slip
	filePath := filepath.Join(destination, f.Name)
	if !strings.HasPrefix(filePath, filepath.Clean(destination)+string(os.PathSeparator)) {
		return fmt.Errorf("invalid file path: %s", filePath)
	}

	// 5. Create directory tree
	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	// 6. Create a destination file for unzipped content
	destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// 7. Unzip the content of a file and copy it to the destination file
	zippedFile, err := f.Open()
	if err != nil {
		return err
	}
	defer zippedFile.Close()

	if _, err := io.Copy(destinationFile, zippedFile); err != nil {
		return err
	}
	return nil
}
