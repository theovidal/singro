package main

type MacroOutput interface {
	Activate()
}

type MacroConfig map[int]MacroGroup

type MacroGroup struct {
	Macros []Macro
	Delay  int
}

func (group *MacroGroup) Execute() {
	var input KeyboardInput
	input.Type = 1
	for _, macro := range group.Macros {
		for i := -1; i < macro.Repeat; i++ {
			if macro.Shift {
				input.ki.wVk = 0x10
				input.Press()
			}
			if macro.Control {
				input.ki.wVk = 0x11
				input.Press()
			}

			for _, key := range macro.Keys {
				key.Activate()
			}

			if macro.Shift {
				input.ki.wVk = 0x10
				input.Release()
			}
			if macro.Control {
				input.ki.wVk = 0x11
				input.Release()
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
	Delay    int
	Duration int
}

func (k Key) Activate() {
	var kbInput KeyboardInput
	kbInput.Type = 1

	var mouseInput MouseInput

	if k.Right {
		mouseInput.mi.dwFlags |= 0x0008
	}
	if k.Middle {
		mouseInput.mi.dwFlags |= 0x0020
	}
	if k.Left {
		mouseInput.mi.dwFlags |= 0x0002
	}

	Sleep(k.Delay)
	if k.Type == "keyboard" {
		kbInput.ki.wVk = uint16(k.Key)
		kbInput.Press()
	} else {
		mouseInput.Apply()
	}

	mouseInput.mi.dwFlags = 0
	if k.Right {
		mouseInput.mi.dwFlags |= 0x0010
	}
	if k.Middle {
		mouseInput.mi.dwFlags |= 0x0040
	}
	if k.Left {
		mouseInput.mi.dwFlags |= 0x0004
	}

	Sleep(k.Duration)
	if k.Type == "keyboard" {
		kbInput.Release()
	} else {
		mouseInput.Apply()
	}
}
