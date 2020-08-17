package main


func (bp *BasePhys) Collides(collision *BasePhys) bool {
	// Log("COLLISION TEST")
	// Take phys objs and make the map the coordinates to variables
	ax1 := bp.x
	ay1 := bp.y
	ax2 := bp.x+bp.w
	ay2 := bp.y+bp.h
	bx1 := collision.x
	by1 := collision.y
	bx2 := collision.x+collision.w
	by2 := collision.y+collision.h

	if ax1 < bx2 && ax2 > bx1 && ay1 < by2 && ay2 > by1 {
		return true
	}

	return false
	
}
