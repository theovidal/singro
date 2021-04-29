package main

import (
	"log"
	"strconv"
	"time"
)

func main() {
	config, err := Open()
	if err != nil {
		log.Fatalf("‼ Error opening or creating configuration file: %s", err)
	}

	registeredMacros := make(MacroConfig)
	for key, macro := range config.Global {
		registeredMacros[key] = macro
	}
	if len(config.Configs) > 0 {
		if config.Selected > len(config.Configs)-1 || config.Selected < 0 {
			log.Fatalln("‼ Selected configuration index out of range")
		}

		for key, macro := range config.Configs[config.Selected] {
			registeredMacros[key] = macro
		}
	}

	log.Printf("▶ singro now running with %d macros", len(registeredMacros))

	stack := make(map[int]bool)
	for {
		time.Sleep(1 * time.Millisecond)

		for key, macro := range registeredMacros {
			intValue := GetKeyState(key)
			binValue := strconv.FormatInt(int64(intValue), 2)

			if len(binValue) < 2 {
				stack[key] = false
				continue
			}

			if !stack[key] {
				log.Printf("ℹ Macro on key %d executed", key)
				macro.Execute()
				stack[key] = true
			}
		}
	}
}
