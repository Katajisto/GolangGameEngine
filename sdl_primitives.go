package main

/*
    SDL DRAWING PRIMITIVES 
*/

import(
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"

)

const(
	windowTitle="Engine";
	windowWidth=1000;
	windowHeight=1000;
)

var SDL_WINDOW_POINTER *sdl.Window
var SDL_SURFACE_POINTER *sdl.Surface
var SDL_RENDERER *sdl.Renderer

func ToggleFullscreen() {
	fullscreenFlag := uint32(sdl.WINDOW_FULLSCREEN)
	isFullscreen := SDL_WINDOW_POINTER.GetFlags() & fullscreenFlag
	
	if isFullscreen == 0 {
		dm, err := sdl.GetCurrentDisplayMode(0)
		if err != nil {
			Log("Error getting display mode: ", err)
		}
		Log(dm.W, dm.H)
		SDL_WINDOW_POINTER.SetFullscreen(fullscreenFlag)
		sdl.ShowCursor(0)
		SDL_WINDOW_POINTER.SetSize(dm.W, dm.H)
		cameraW = float64(dm.W)
		cameraH = float64(dm.H)
		
	} else {
		SDL_WINDOW_POINTER.SetFullscreen(0)
		sdl.ShowCursor(1)
		SDL_WINDOW_POINTER.SetSize(1000, 1000)
	}
	
}

func sdlStart() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	if err := ttf.Init(); err != nil {
		Log("Failed to initialize TTF: %s\n", err)
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
	SDL_RENDERER, err = sdl.CreateRenderer(SDL_WINDOW_POINTER, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		Log("Failed to create renderer: %s\n", err)
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
	SDL_RENDERER.Present()
}

func Clear() {
	SDL_RENDERER.SetDrawColor(0,0,0,255)
	SDL_RENDERER.Clear()
}

func DrawRect(y,x,w,h float64) {
	y = cameraY - y
	x = x - cameraX
	rect := sdl.Rect{FtoI(x),FtoI(y),FtoI(w),FtoI(h)}
	SDL_RENDERER.SetDrawColor(0,0,255,255)
	SDL_RENDERER.FillRect(&rect)
}

func LoadFont(fontName string, size int) *ttf.Font {
	var font *ttf.Font
	var err error
	if font, err = ttf.OpenFont(fontName, size); err != nil {
		Log("Failed to open font: %s\n", err)
		return nil
	}
	return font
}

var DEBUG_FONT *ttf.Font

func CreateFontSurface(font *ttf.Font, text string) *sdl.Surface  {
	var err error
	var solidSurface *sdl.Surface
	if solidSurface, err = font.RenderUTF8Solid(text, sdl.Color{255,255,255,255}); err != nil {
		Log("Failed to create font surface")
	}
	return solidSurface
}

func CreateFontTexture(font *ttf.Font, text string) *sdl.Texture {
	var texture *sdl.Texture
	var err error
	surface := CreateFontSurface(font, text)
	if texture, err = SDL_RENDERER.CreateTextureFromSurface(surface); err != nil {
		Log("Failed to create texture")
	}
	return texture
}

func DrawFont(texture *sdl.Texture, x,y,w,h float64) {
	SDL_RENDERER.Copy(texture, nil, &sdl.Rect{FtoI(x),FtoI(y),FtoI(w),FtoI(h)})
}

