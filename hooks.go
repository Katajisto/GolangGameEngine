package main

var drawHook func() = func() {}
var physHook func() = func() {}

func SetDrawHook(hook func()) {
	drawHook = hook
}

func SetPhysHook(hook func()) {
	physHook = hook
}
