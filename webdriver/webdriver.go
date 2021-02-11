/*
 * File: webdriver.go
 * Created: Sunday, 7th February 2021 4:30:08 pm
 * Author: Rakibul Yeasin (ryeasin03@gmail.com)
 * -----
 * Last Modified: Friday, 12th February 2021 3:01:02 am
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
	"path"
	"strconv"
	"syscall"

	"github.com/imroc/req"
)

// Server Details
type Server struct {
	uri       string
	location  string
	pid       int
	sessionId string
	url       string
}

var srv Server
var res map[string]interface{}

// RunServer method
func RunServer() {
	cwdPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	srv.location = path.Join(cwdPath, "webdriver/driver/geckodriver")
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
	srv.sessionId = r["sessionId"].(string)
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
func Kill() {
	err := syscall.Kill(srv.pid, syscall.SIGKILL)
	if err != nil {
		kill()
	}
}

func kill() error {
	kill := exec.Command("TASKKILL", "/T", "/F", "/PID", strconv.Itoa(srv.pid))
	kill.Stderr = os.Stderr
	kill.Stdout = os.Stdout
	return kill.Run()
}
