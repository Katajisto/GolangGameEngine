// +build debug

package main

const(
	DEBUG_MODE = true
)

func initDebug() {
	DEBUG_FONT = LoadFont("ubuntu.ttf", 100)
}

func printDebugData(debug []string) {
	for i, text := range debug {
		texture := CreateFontTexture(DEBUG_FONT, text)
		DrawFont(texture, float64(10), float64(i*22), float64(10*len(text)), 20.0)
	}
}
