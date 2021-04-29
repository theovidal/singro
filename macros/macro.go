package main

type SingroConfig struct {
	Global   MacroConfig
	Selected int
	Configs  []MacroConfig
}

type MacroConfig map[int]MacroGroup

type MacroGroup struct {
	Macros []Macro
	Delay  int
}

func (group *MacroGroup) Execute() {
	for _, macro := range group.Macros {
		iterations := macro.Repeat
		if iterations == 0 { iterations++ }

		for i := 0; i < iterations; i++ {
			if macro.Shift {
				KeyboardPress(0x10)
			}
			if macro.Control {
				KeyboardPress(0x11)
			}

			for _, key := range macro.Keys {
				key.Activate()
			}

			if macro.Shift {
				KeyboardRelease(0x10)
			}
			if macro.Control {
				KeyboardRelease(0x11)
			}

			Sleep(group.Delay)
		}
	}
}

type Macro struct {
	Keys    []Key
	Repeat  int
	Control bool
	Shift   bool
}

type Key struct {
	Type     string
	Repeat   int
	Delay    int
	Duration int

	// KEYBOARD TYPE
	Key int

	// SEQUENCE TYPE
	Keys []int

	// MOUSE TYPE
	Right    bool
	Middle   bool
	Left     bool
	Wheel    uint32
	HWheel   uint32
	X        int32
	Y        int32
	Absolute bool
}

func (k Key) Activate() {
	iterations := k.Repeat
	if iterations == 0 { iterations++ }

	for i := 0; i < iterations; i++ {
		Sleep(k.Delay)
		pressSequences[k.Type](&k)

		Sleep(k.Duration)
		releaseSequences[k.Type](&k)
	}
}
