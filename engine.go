package main

func EngineStart() {
	sdlStart()
	KeepLooping := true
	for KeepLooping {
		KeepLooping = HandleEvent()
		physHook(1)
		drawHook()
		Draw()
		Clear()
		Delay(10)
	}
}
