package main

import (
	"syscall"
	"time"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procGetAsyncKeyState = user32.NewProc("GetKeyState")
)

func Sleep(amount int) {
	time.Sleep(time.Duration(amount*1000) * time.Microsecond)
}
