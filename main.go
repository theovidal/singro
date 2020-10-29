package main

import (
	"strconv"
	"syscall"
	"time"

	"github.com/micmonay/keybd_event"
)

type MacroGroup struct {
	Macros []Macro
	Delay  int
}

type Macro struct {
	Keys    []int
	Control bool
	Shift   bool
}

func (group *MacroGroup) Execute() {
	for _, macro := range group.Macros {
		kb, err := keybd_event.NewKeyBonding()
		if err != nil {
			panic(err)
		}

		kb.SetKeys(macro.Keys...)
		kb.HasSHIFT(macro.Shift)
		kb.HasCTRL(macro.Control)

		err = kb.Launching()
		if err != nil {
			panic(err)
		}

		time.Sleep(time.Duration(group.Delay))
	}

}

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procGetAsyncKeyState = user32.NewProc("GetKeyState")

	macros = map[int]MacroGroup{
		6: {
			Macros: []Macro{
				{
					Keys: []int{keybd_event.VK_T, keybd_event.VK_G},
				},
				{
					Keys: []int{keybd_event.VK_G, keybd_event.VK_ENTER},
				},
			},
		},
		7: {
			Macros: []Macro{
				{
					Keys: []int{keybd_event.VK_F5},
				},
			},
		},
	}
)

func GetKeyState(key int) uintptr {
	value, _, _ := procGetAsyncKeyState.Call(uintptr(key))
	return value
}

func main() {
	for {
		time.Sleep(1 * time.Millisecond)
		for KEY := 0; KEY <= 256; KEY++ {
			pressed := false
			for {
				intValue := GetKeyState(KEY)
				binValue := strconv.FormatInt(int64(intValue), 2)

				if len(binValue) < 2 {
					break
				}

				if !pressed {
					macro, found := macros[KEY]
					if !found {
						continue
					}

					macro.Execute()
				}

				pressed = true
			}
		}
	}
}
