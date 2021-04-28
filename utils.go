package main

import (
	"time"
)

func Sleep(amount int) {
	time.Sleep(time.Duration(amount*1000) * time.Microsecond)
}
