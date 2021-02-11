/*
 * File: main.go
 * Created: Sunday, 7th February 2021 4:31:07 pm
 * Author: Rakibul Yeasin (ryeasin03@gmail.com)
 * -----
 * Last Modified: Friday, 12th February 2021 2:45:15 am
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
	driver.Get(`https://google.com`)
	driver.GetStatus()

	time.Sleep(8 * time.Second)
	driver.GetStatus()
	driver.Screenshot("google")
	time.Sleep(8 * time.Second)

	driver.Kill()
}
