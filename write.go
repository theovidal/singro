package main

import (
	"log"
	"unsafe"
)

type KeyboardInput struct {
	Type uint32
	ki   struct {
		wVk         uint16
		wScan       uint16
		dwFlags     uint32
		time        uint32
		dwExtraInfo uint64
	}
	Size uint64
}

func (i KeyboardInput) Apply() {
	_, _, output := procSendInput.Call(
		uintptr(1),
		uintptr(unsafe.Pointer(&i)),
		unsafe.Sizeof(i),
	)
	log.Println(output)
}

func (i *KeyboardInput) Press() {
	i.ki.dwFlags = 0
	i.Apply()
}

func (i *KeyboardInput) Release() {
	i.ki.dwFlags |= 0x0002
	i.Apply()
	i.ki.dwFlags = 0
}

// UNUSED CODE - Windows API doesn't want to get it, so we'll pass through C

type MouseInput struct {
	Type uint32
	mi   struct {
		dx          int32
		dy          int32
		mouseData   uint32
		dwFlags     uint32
		time        uint32
		dwExtraInfo uint64
	}
	Size uint64
}

func (i MouseInput) Apply() {
	_, _, output := procSendInput.Call(
		uintptr(1),
		uintptr(unsafe.Pointer(&i)),
		unsafe.Sizeof(i),
	)
	log.Println(output)
}
