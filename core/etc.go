package core

// Vec represents a 2d Vector.
type Vec struct {
	X int
	Y int
}

func (v Vec) add(v2 Vec) Vec {
	return Vec{v.X + v2.X, v.Y + v2.Y}
}
func (v Vec) sub(v2 Vec) Vec {
	return Vec{v.X - v2.X, v.Y - v2.Y}
}
func (v Vec) mult(v2 Vec) Vec {
	return Vec{v.X * v2.X, v.Y * v2.Y}
}
func (v Vec) div(v2 Vec) Vec {
	return Vec{v.X / v2.X, v.Y / v2.Y}
}
