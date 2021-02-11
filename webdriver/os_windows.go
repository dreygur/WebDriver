/*
 * File: os_windows.go
 * Created: Friday, 12th February 2021 3:31:46 am
 * Author: Rakibul Yeasin (ryeasin03@gmail.com)
 * -----
 * Last Modified: Friday, 12th February 2021 3:33:44 am
 * Modified By: Rakibul Yeasin (ryeasin03@gmail.com)
 * -----
 * Copyright (c) 2021 Slishee
 */

package webdriver

import (
	"os"
	"os/exec"
	"strconv"
)

// Kill the Server
func Kill() error {
	kill := exec.Command("TASKKILL", "/T", "/F", "/PID", strconv.Itoa(srv.pid))
	kill.Stderr = os.Stderr
	kill.Stdout = os.Stdout
	return kill.Run()
}
