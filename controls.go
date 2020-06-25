package main

var keyStateMap map[string]bool
var keyHookMap map[string][]func()
var keyPressHookMap map[string][]func()

func init() {
	keyStateMap = make(map[string]bool)
	keyHookMap = make(map[string][]func())
	keyPressHookMap = make(map[string][]func())
}

//Receives a bool wheter the key is down or not.
func DispatchKey(keyname string, state bool) {
	// Fire key 
	for _, hook := range keyHookMap[keyname] {
		hook()
	}
	// Update key state
	keyStateMap[keyname] = state
}

func GetKeyState(keyname string) bool {
	return keyStateMap[keyname]
}

func AddKeyHook(keyname string, hook func()) {
	keyHookMap[keyname] = append(keyHookMap[keyname], hook)
}

func AddKeyPressHook(keyname string, hook func()) {
	keyPressHookMap[keyname] = append(keyPressHookMap[keyname], hook)
}
