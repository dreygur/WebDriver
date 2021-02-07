/*
 * File: main.go
 * Created: Sunday, 7th February 2021 4:31:07 pm
 * Author: Rakibul Yeasin (ryeasin03@gmail.com)
 * -----
 * Last Modified: Sunday, 7th February 2021 6:03:43 pm
 * Modified By: Rakibul Yeasin (ryeasin03@gmail.com)
 * -----
 * Copyright (c) 2021 Slishee
 */

package main

import (
	"time"
	driver "webdriver/webdriver"
)

func main() {
	driver.RunServer()
	_ = driver.GetSession()
	// fmt.Println(res)
	time.Sleep(8 * time.Second)
	driver.Kill()
}
