package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"syscall"
	"time"

	"gopkg.in/yaml.v2"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procGetAsyncKeyState = user32.NewProc("GetKeyState")
	procSendInput        = user32.NewProc("SendInput")
)

func main() {
	path := "./macros.yml"

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln(err)
	}

	var macros MacroConfig
	err = yaml.Unmarshal(data, &macros)
	if err != nil {
		log.Panicln(err)
	}

	log.Println("singro started")

	stack := make(map[int]bool)
	for {
		time.Sleep(1 * time.Millisecond)

		for key, macro := range macros {
			intValue := GetKeyState(key)
			binValue := strconv.FormatInt(int64(intValue), 2)

			if len(binValue) < 2 {
				stack[key] = false
				continue
			}

			if !stack[key] {
				log.Printf("Macro on key %d executed", key)
				macro.Execute()
				stack[key] = true
			}
		}
	}
}
