package main

var keyStateMap map[string]bool

func init() {
	keyStateMap = make(map[string]bool)
}

//Receives a bool wheter the key is down or not.
func DispatchKey(keyname string, state bool) {
	keyStateMap[keyname] = state
}

func GetKeyState(keyname string) bool {
	return keyStateMap[keyname]
}
