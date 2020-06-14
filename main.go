package main

// Just a test driver function

var x int32 = 0

func DrawFunc() {
	if(GetKeyState("A")) {
		x += 10;
	} else {
		x++;
	}
	if x > 800 { 
		DrawRect(x,0,200,200)
		DrawRect((x-800)-200,0,200,200)
	} else {
		DrawRect(x,0,200,200)
	}

	if x > 1000 {
		x = 0
	}
}

func main() {
	SetDrawHook(DrawFunc)
	EngineStart ()
	CleanUp()
	
}
