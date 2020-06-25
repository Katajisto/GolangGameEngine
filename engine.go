package main

import(
	"time"
	"strconv"
)

var avg_draw_time int64 = 5
var avg_phys_time int64 = 5

const(
	PHYS_FRAME_SIZE = 10000
)

func EngineInit() {
	sdlStart()
	initDebug()
}

func EngineStart() {
	KeepLooping := true
	runtime := time.Now()
	frames := 0
	frameTimeCount := int64(0)
	for KeepLooping {
		KeepLooping = HandleEvent()
		frameStartTime := time.Now()
		Clear()
		if DEBUG_MODE {
			printDebugData([]string{"AVG DRAW: " + strconv.Itoa(avg_draw_time), "AVG PHYS: " + strconv.Itoa(avg_phys_time), "AVERAGE FPS: ", strconv.Itoa(frames/int(time.Now().Sub(runtime).Seconds()))})
		}
		drawHook()
		DrawCameraArea()
		Draw()
		frameEndTime := time.Now()
		frameTime := frameEndTime.Sub(frameStartTime)
		frameTimeCount += frameTime.Microseconds()
		avg_draw_time = (avg_draw_time+frameTime.Microseconds())/2
		for frameTimeCount >= PHYS_FRAME_SIZE {
			physStartTime := time.Now()
			physHook()
			calcEntityPhys()
			physEndTime := time.Now()
			avg_phys_time = (avg_phys_time+physEndTime.Sub(physStartTime).Microseconds())/2
			frameTimeCount -= PHYS_FRAME_SIZE
		}
		frames++
	}
	Log("AVERAGE FPS: ", (frames/int(time.Now().Sub(runtime).Seconds())))
}

func PrintStat() {
	Log(avg_draw_time, " :: ", avg_phys_time)
}
