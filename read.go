package main

func GetKeyState(key int) uintptr {
	value, _, _ := procGetAsyncKeyState.Call(uintptr(key))
	return value
}
