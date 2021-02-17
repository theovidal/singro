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
	var input Input
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
	Right    bool
	Middle   bool
	Left     bool
	Key      int
	Delay    int
	Duration int
}

func (k Key) Activate() {
	var input Input
	input.Type = 1

	input.ki.wVk = uint16(k.Key)

	/*if k.Right {
		input.mi.dwFlags |= 0x0008
	}
	if k.Middle {
		input.mi.dwFlags |= 0x0020
	}
	if k.Left {
		input.mi.dwFlags |= 0x0002
	}*/

	Sleep(k.Delay)
	input.Press()

	/*input.mi.dwFlags = 0
	if k.Right {
		input.mi.dwFlags |= 0x0010
	}
	if k.Middle {
		input.mi.dwFlags |= 0x0040
	}
	if k.Left {
		input.mi.dwFlags |= 0x0004
	}*/

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
