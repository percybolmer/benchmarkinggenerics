package benchmarking

// Below is the Type casting based Solution
//
type CarRegular struct {
	Name          string
	DistanceMoved int
}

type PersonRegular struct {
	Name          string
	DistanceMoved float32
}

func (p *PersonRegular) Move(meters float32) {
	p.DistanceMoved += meters
}

func (c *CarRegular) Move(meters int) {
	c.DistanceMoved += meters
}

func MoveRegular(v interface{}, distance float32) {
	switch v.(type) {
	case *PersonRegular:
		v.(*PersonRegular).Move(distance)
	case *CarRegular:
		v.(*CarRegular).Move(int(distance))
	default:
		// Handle Unsupported types, not needed by Generic solution as Compiler does this for you
	}
}
