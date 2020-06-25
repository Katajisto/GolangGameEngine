package main

import "math"

//Basic vector type
type Vector struct {
	x float64
	y float64
}


//Transforms float into int32 and rounds it mathematically.
func FtoI(f float64) int32 {
	return(int32(math.Round(f)))
}

//Adds argument vector to the original vector
func (v1 Vector) Add(v2 Vector) Vector {
	v1.x += v2.x;
	v1.y += v2.y
	return v1
}

//Subtracts argument vector from original vector
func (v1 Vector) Sub(v2 Vector) Vector {
	v1.x -= v2.x;
	v1.y -= v2.y;
	return v1
}

func (v1 Vector) Mul(n float64) Vector{
	v1.x *= n
	v1.y *= n
	return v1
}

func VAbs(v Vector) Vector {
	v.x = math.Abs(v.x)
	v.y = math.Abs(v.y)
	return v
}
