package main

var drawHook func() = func() {}
var physHook func(int64) = func(int64) {}

func SetDrawHook(hook func()) {
	drawHook = hook
}

func SetPhysHook(hook func(int64)) {
	physHook = hook
}
