package main

/*
    SDL DRAWING PRIMITIVES 
*/

import(
	"github.com/veandco/go-sdl2/sdl"
)

const(
	windowTitle="Engine";
	windowWidth=1000;
	windowHeight=1000;
)

var SDL_WINDOW_POINTER *sdl.Window
var SDL_SURFACE_POINTER *sdl.Surface



func Start() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	var err error
	SDL_WINDOW_POINTER, err = sdl.CreateWindow(windowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, windowHeight, windowWidth, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	SDL_SURFACE_POINTER, err = SDL_WINDOW_POINTER.GetSurface()
	if err != nil {
		panic(err)
	}
}

func CleanUp() {
	SDL_WINDOW_POINTER.Destroy()
}

func HandleEvent() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			return false
			break
		case *sdl.KeyboardEvent:
			if event.GetType() == sdl.KEYDOWN {
				DispatchKey(sdl.GetKeyName(t.Keysym.Sym), true)
			} else if event.GetType() == sdl.KEYUP {
				DispatchKey(sdl.GetKeyName(t.Keysym.Sym), false)
			}
		}
	}
	return true
}

func Draw() {
	SDL_WINDOW_POINTER.UpdateSurface()
}

func Clear() {
	SDL_SURFACE_POINTER.FillRect(nil,0)
}

func DrawRect(y,x,w,h int32) {
	rect := sdl.Rect{y,x,w,h}
	SDL_SURFACE_POINTER.FillRect(&rect, 0xffff0000)
}
