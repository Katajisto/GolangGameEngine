package main

type BasePhys struct {
	x float64
	y float64
	w float64
	h float64
	weight float64
	velocity Vector
	friction float64
}


//Applies a 2D vector to a base phys object.
func (base *BasePhys) Apply(force Vector) {
	base.velocity = base.velocity.Add(force)
}

//Takes a vector and subtracts its value from the basePhys
func (base *BasePhys) Brake(force Vector) {
	force = VAbs(force)
	velocity := &base.velocity
	if velocity.x > 0 {
		velocity.x -= force.x
		if velocity.x < 0 {
			velocity.x = 0
		}
	} else if velocity.x < 0 {
		velocity.x += force.x
		if velocity.x > 0 {
			velocity.x = 0
		}
	}	
	if velocity.y > 0 {
		velocity.y -= force.y
		if velocity.y < 0 {
			velocity.y = 0
		}
	} else if velocity.y < 0 {
		velocity.y += force.y
		if velocity.y > 0 {
			velocity.y = 0
		}
	}
}

func (b *BasePhys) phys() {
	//Apply gravity
	b.Apply(Vector{0,-0.1}.Mul(b.weight))
	b.x += b.velocity.x
	b.y += b.velocity.y
	// TODO: isTouchingGround
	if b.y <= 100 {
		b.Brake(Vector{0.01*b.friction,0})
	}
	if b.y <= 100 {
		b.y = 100;
		b.velocity.y = 0;
	}
}

func (b *BasePhys) setVelocity(newVelocity Vector) {
	b.velocity = newVelocity
}

func (b *BasePhys) setYVelocity(new float64) {
	b.velocity.y = new
}

func (b *BasePhys) setXVelocity(new float64) {
	b.velocity.x = new
}
