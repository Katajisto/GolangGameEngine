package main

// Just a test driver function

var x int32 = 0

type Block struct {
	phys BasePhys
}

func (block *Block) Draw() {
	b := block.phys
	DrawRect(b.y, b.x, b.w, b.h)
}

func (block *Block) GetBasePhys() *BasePhys {
	return &block.phys
}

func main() {
	InitCamera(0,1000,1000,1000)
	test := &Block{BasePhys{0,100,100,100,5,Vector{0,0},25}}
	collider := &Block{BasePhys{200,300,300,100,0,Vector{0,0},25}}
	collider2 := &Block{BasePhys{400,700,300,100,0,Vector{0,0},25}}
	AddEntity(test)
	AddEntity(collider)
	AddEntity(collider2)
	AddKeyHook("Space", func() {
		test.phys.setYVelocity(20)
 	})
	AddKeyHook("F", ToggleFullscreen)
	SetPhysHook(func() {
		if GetKeyState("A") {
			test.phys.Apply(Vector{-0.4,0})
		}
		if GetKeyState("D") {
			test.phys.Apply(Vector{0.4,0})
		}
		if test.phys.velocity.x > 6 { 
			test.phys.velocity.x = 6
		}
		if test.phys.velocity.x < -6 {
			test.phys.velocity.x = -6
		}
		if collider.phys.Collides(&test.phys) {
			Log("COLLISION")
		}
		//		SetCameraPos(test.phys.x-500, 1000)
	})
	AddKeyPressHook("D", func() {
		test.phys.setXVelocity(3)
	})
	AddKeyPressHook("A", func() {
		test.phys.setXVelocity(-3)
	})
	EngineInit()
	SetDrawHook(func() {
		//	DrawFont(texture)
	})
	EngineStart ()
	CleanUp()
	PrintStat()
	
}
