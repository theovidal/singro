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
		for i := -1; i < macro.Repeat; i++ {
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
	Right    bool
	Middle   bool
	Left     bool
	Key      int
	Keys     []int
	Delay    int
	Duration int
}

func (k Key) Activate() {
	Sleep(k.Delay)
	pressSequences[k.Type](&k)

	Sleep(k.Duration)
	releaseSequences[k.Type](&k)
}
