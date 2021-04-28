package main

/*
#include <Windows.h>

void mouseInput(DWORD flags, DWORD mouseData, LONG dx, LONG dy) {
  INPUT input;
  input.type = INPUT_MOUSE;

  input.mi.dx          = dx;
  input.mi.dy          = dy;
  input.mi.dwFlags     = flags;
  input.mi.mouseData   = mouseData;
  input.mi.dwExtraInfo = 0;
  input.mi.time        = 0;

	SendInput(1, &input, sizeof(INPUT));
}

void keyboardInput(WORD keyCode, _Bool release) {
	INPUT input;
  input.type = INPUT_KEYBOARD;

  input.ki.wVk         = keyCode;
	input.ki.wScan       = 0;
  input.ki.time        = 0;
  input.ki.dwExtraInfo = 0;
	if (release)
  	input.ki.dwFlags   = 2;
  else
		input.ki.dwFlags   = 0;

  SendInput(1,&input,sizeof(INPUT));
}
*/
import "C"

var pressSequences = map[string]func(k *Key){
	"keyboard": func(k *Key) {
		KeyboardPress(k.Key)
	},
	"sequence": func(k *Key) {
		for _, key := range k.Keys {
			KeyboardPress(key)
			KeyboardRelease(key)
		}
	},
	"mouse": func(k *Key) {
		var flags uint32
		var mouseData uint32
		var dx int32
		var dy int32

		if k.Right {
			flags |= 0x0008
		}
		if k.Middle {
			flags |= 0x0020
		}
		if k.Left {
			flags |= 0x0002
		}

		if k.X != 0 || k.Y != 0 || k.Absolute {
			flags |= 0x0001
			if k.Absolute {
				flags |= 0x8000
			}
			dx = k.X
			dy = k.Y
		}

		if k.Wheel != 0 {
			flags |= 0x0800
			mouseData = k.Wheel
		}
		if k.HWheel != 0 {
			flags |= 0x1000
			mouseData = k.HWheel
		}

		C.mouseInput((C.ulong)(flags), (C.ulong)(mouseData), (C.long)(dx), (C.long)(dy))
	},
}

var releaseSequences = map[string]func(k *Key){
	"keyboard": func(k *Key) {
		KeyboardRelease(k.Key)
	},
	"sequence": func(_ *Key) {},
	"mouse": func(k *Key) {
		var mouseInput C.ulong
		if k.Right {
			mouseInput |= 0x0010
		}
		if k.Middle {
			mouseInput |= 0x0040
		}
		if k.Left {
			mouseInput |= 0x0004
		}

		C.mouseInput(mouseInput, 0, 0, 0)
	},
}

func KeyboardPress(key int) {
	C.keyboardInput((C.ushort)(key), false)
}
func KeyboardRelease(key int) {
	C.keyboardInput((C.ushort)(key), true)
}
