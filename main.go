package main

import (
	"log"
	"strconv"
	"syscall"
	"time"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procGetAsyncKeyState = user32.NewProc("GetKeyState")
	procSendInput        = user32.NewProc("SendInput")

	macros = map[int]MacroGroup{
		6: {
			Macros: []Macro{
				{
					Keys: []MacroOutput{
						Key{Key: 0x54},
						Key{Key: 0x47, Delay: 50},
						Key{Key: 0x47},
						Key{Key: 0x0D},
					},
					Shift: true,
				},
			},
		},
		0xA3: {
			Macros: []Macro{
				{
					Keys: []MacroOutput{
						Mouse{Left: true},
					},
				},
			},
		},
		5: {
			Macros: []Macro{
				{
					Keys: []MacroOutput{
						Key{Key: 0x74},
					},
				},
			},
		},
	}
)

func main() {
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
