package main

type MacroGroup struct {
	Macros []Macro
	Delay  int
}

type Macro struct {
	Keys    []MacroOutput
	Repeat  int
	Control bool
	Shift   bool
}

type Key struct {
	Key      int
	Delay    int
	Duration int
}

func (k Key) Activate() {
	var input KeyboardInput
	input.Type = 1

	input.ki.wVk = uint16(k.Key)
	Sleep(k.Delay)
	input.Press()
	Sleep(k.Duration)
	input.Release()
}

type Mouse struct {
	Right  bool
	Middle bool
	Left   bool
}

func (m Mouse) Activate() {
	var input MouseInput

	if m.Right {
		input.dwFlags |= 0x0008
	}
	if m.Middle {
		input.dwFlags |= 0x0020
	}
	if m.Left {
		input.dwFlags |= 0x0002
	}

	input.Apply()

	input.dwFlags = 0
	if m.Right {
		input.dwFlags |= 0x0010
	}
	if m.Middle {
		input.dwFlags |= 0x0040
	}
	if m.Left {
		input.dwFlags |= 0x0004
	}

	input.Apply()
}

type MacroOutput interface {
	Activate()
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
