/*
 * File: os_unix.go
 * Created: Friday, 12th February 2021 3:31:53 am
 * Author: Rakibul Yeasin (ryeasin03@gmail.com)
 * -----
 * Last Modified: Friday, 12th February 2021 3:33:29 am
 * Modified By: Rakibul Yeasin (ryeasin03@gmail.com)
 * -----
 * Copyright (c) 2021 Slishee
 */

package webdriver

import "syscall"

// Kill the Server
func Kill() error {
	return syscall.Kill(srv.pid, syscall.SIGKILL)
}
