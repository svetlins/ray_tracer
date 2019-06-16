package common

import (
	"math"
)

type Pattern struct {
	colors    []Vector
	transform Matrix
	inverseTransform Matrix
	value     func(p Vector, colors []Vector) Vector
}

func PatternColors(colors ...Vector) func(*Pattern) {
	return func (p *Pattern) {
		p.colors = colors
	}
}

func PatternTransform(transform ...Matrix) func(*Pattern) {
	return func (p *Pattern) {
		p.transform = Transform(transform...)
		p.inverseTransform = p.transform.Inverse()
	}
}

func PatternValue(value func(Vector, []Vector) Vector) func(*Pattern) {
	return func (p *Pattern) {
		p.value = value
	}
}

func NewPattern(options ...func(*Pattern)) *Pattern {
	pattern := &Pattern{
		colors: []Vector{Black(), White()},
		transform: NewIdentity4(),
		value:     Stripe(),
	}

	for _, option := range options {
		option(pattern)
	}

	return pattern
}

func (s *Pattern) SetTransform(transform Matrix) {
	s.transform = transform
}

func (s *Pattern) SetValueFunc(value func(p Vector, colors []Vector) Vector) {
	s.value = value
}

func (s Pattern) At(point Vector) Vector {
	return s.value(point, s.colors)
}

func (s Pattern) AtObject(o Object, point Vector) Vector {
	point = o.transformPointToObject(point)
	point = s.inverseTransform.Transform(point)

	return s.At(point)
}

func TestPattern() func(p Vector, colors []Vector) Vector {
	return func(p Vector, colors []Vector) Vector {
		return NewColor(p.x, p.y, p.z)
	}
}

func Stripe() func(p Vector, colors []Vector) Vector {
	return func(p Vector, colors []Vector) Vector {
		if int(math.Floor(p.x*5))%2 == 0 {
			return colors[0]
		} else {
			return colors[1]
		}
	}
}

func Gradient() func(p Vector, colors []Vector) Vector {
	return func(p Vector, colors []Vector) Vector {
		c1 := colors[0]
		c2 := colors[1]

		distance := c2.Subtract(c1)
		fraction := math.Abs(p.x - math.Floor(p.x) - 0.5)

		return c1.Add(distance.Scale(fraction * 2))
	}
}

func Ring() func(p Vector, colors []Vector) Vector {
	return func(p Vector, colors []Vector) Vector {
		if int(math.Floor(math.Sqrt(p.x*p.x+p.z*p.z)))%2 == 0 {
			return colors[0]
		} else {
			return colors[1]
		}
	}
}

func Checkers() func(p Vector, colors []Vector) Vector {
	return func(p Vector, colors []Vector) Vector {
		x := int(math.Floor(p.x))
		y := int(math.Floor(p.y))
		z := int(math.Floor(p.z))

		if (x+y+z)%2 == 0 {
			return colors[0]
		} else {
			return colors[1]
		}
	}
}
