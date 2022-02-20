package benchmarking



// Subtractable is a type constraint that defines subtractable datatypes to be used in generic functions
type Subtractable interface {
	int | int64 | float32
}
// Moveable is the interace for moving a Entity
type Moveable[S Subtractable] interface {
	Move(S)
}

// Car is a Generic Struct with the type S to be defined
type Car[S Subtractable] struct {
	Name string
	DistanceMoved S
}

// Person is a Generic Struct with the type S to be defined
type Person[S Subtractable] struct {
	Name string
	DistanceMoved S
}

// Person is a struct that accepts a type definition at initialization
// And uses that Type as the data type for meters as input
func (p *Person[S]) Move(meters S) {
	p.DistanceMoved += meters
}
func (c *Car[S]) Move(meters S) {
	c.DistanceMoved += meters
}

// Move is a generic function that takes in a Generic Moveable and moves it
func Move[S Subtractable, V Moveable[S]](v V, meters S) {
	v.Move(meters)
}

