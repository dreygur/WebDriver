/*
 * File: webdriver.go
 * Created: Sunday, 7th February 2021 4:30:08 pm
 * Author: Rakibul Yeasin (ryeasin03@gmail.com)
 * -----
 * Last Modified: Friday, 12th February 2021 4:42:22 am
 * Modified By: Rakibul Yeasin (ryeasin03@gmail.com)
 * -----
 * Copyright (c) 2021 Slishee
 */

package webdriver

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/imroc/req"
)

// Server Details
type Server struct {
	uri       string
	location  string
	pid       int
	sessionID string
	url       string
}

var srv Server
var res map[string]interface{}

// RunServer method
func RunServer(webdriverPath string) {
	cwdPath, err := os.Getwd()
	fmt.Println(cwdPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// if runtime.GOOS == "linux" {
	// 	srv.location = path.Join(cwdPath, "webdriver", "driver", "geckodriver")
	// } else if runtime.GOOS == "windows" {
	// 	srv.location = path.Join(cwdPath, "webdriver", "driver", "geckodriver.exe")
	// }
	srv.location = webdriverPath

	gecko := exec.Command(srv.location)
	err = gecko.Start()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	srv.pid = gecko.Process.Pid
}

// GetSession returns session data
func GetSession() interface{} {
	srv.uri = "http://localhost:4444/session/"
	// Headers
	header := req.Header{
		"Content-Type": "application/json",
	}

	// Body
	body, err := json.Marshal(map[string]interface{}{
		"capabilities": map[string]interface{}{
			"alwaysMatch": map[string]interface{}{
				"acceptInsecureCerts": true,
			},
		},
	})
	if err != nil {
		return err.Error()
	}

	resp, err := req.Post(srv.uri, header, body)
	if err != nil {
		return err.Error()
	}
	resp.ToJSON(&res)
	r := res["value"].(map[string]interface{})
	srv.sessionID = r["sessionId"].(string)
	// fmt.Println(resp)
	return resp
}

// GetStatus of current session
func GetStatus() interface{} {
	// Headers
	header := req.Header{
		"Content-Type": "application/json",
	}

	resp, err := req.Get(`https://localhost:4444/session/honululu/google.com`, header)
	if err != nil {
		fmt.Println((err))
		return err.Error()
	}
	resp.ToJSON(&res)
	fmt.Println(resp)
	return resp
}

// Kill the Server
func Kill() error {
	var err error
	// err := syscall.Kill(srv.pid, syscall.SIGKILL)
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		k := exec.Command("kill", "-9", strconv.Itoa(srv.pid))
		k.Stderr = os.Stderr
		k.Stdout = os.Stdout
		err = k.Run()
	} else {
		k := exec.Command("TASKKILL", "/T", "/F", "/PID", strconv.Itoa(srv.pid))
		k.Stderr = os.Stderr
		k.Stdout = os.Stdout
		err = k.Run()
	}
	return err
}
