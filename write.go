package main

/*
#include <Windows.h>

void mouseInput(DWORD flags) {
  INPUT input;
  input.type = INPUT_MOUSE;

  input.mi.dx          = 0;
  input.mi.dy          = 0;
  input.mi.dwFlags     = flags | 1;
  input.mi.mouseData   = 0;
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
	"mouse": func(k *Key) {
		var mouseInput C.ulong

		if k.Right {
			mouseInput |= 0x0008
		}
		if k.Middle {
			mouseInput |= 0x0020
		}
		if k.Left {
			mouseInput |= 0x0002
		}

		C.mouseInput(mouseInput)
	},
}

var releaseSequences = map[string]func(k *Key){
	"keyboard": func(k *Key) {
		KeyboardRelease(k.Key)
	},
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

		C.mouseInput(mouseInput)
	},
}

func KeyboardPress(key int) {
	C.keyboardInput((C.ushort)(key), false)
}
func KeyboardRelease(key int) {
	C.keyboardInput((C.ushort)(key), true)
}
