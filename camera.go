package main

var cameraX float64
var cameraY float64
var cameraW float64
var cameraH float64

func InitCamera(x,y,w,h float64) {
	cameraX = x
	cameraY = y
	cameraW = w
	cameraH = h
}

func SetCameraPos(x, y float64) {
	cameraX = x
	cameraY = y
}

func DrawCameraArea() {
	for _, entity := range entityList {
		entityPhys := entity.GetBasePhys()
		if entityPhys.x + entityPhys.w >= cameraX && entityPhys.x <= cameraX + cameraW {
			if entityPhys.y - entityPhys.h <= cameraY && entityPhys.y >= cameraY-cameraH {
				entity.Draw()
			}
		}
	}
}
