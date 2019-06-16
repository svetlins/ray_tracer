package common

type Material struct {
	pattern         *Pattern
	color           Vector
	ambient         float64
	diffuse         float64
	specular        float64
	shininess       float64
	reflective      float64
	transparency    float64
	refractiveIndex float64
}

func MaterialPattern(p *Pattern)  func(*Material) {
	return func (m *Material) {
		m.pattern = p
	}
}

func Color(c Vector) func(*Material) {
	return func (m *Material) {
		m.color = c
	}
}

func Ambient(value float64) func(*Material) {
	return func (m *Material) {
		m.ambient = value
	}
}

func Diffuse(value float64) func(*Material) {
	return func (m *Material) {
		m.diffuse = value
	}
}

func Specular(value float64) func(*Material) {
	return func (m *Material) {
		m.specular = value
	}
}

func Shininess(value float64) func(*Material) {
	return func (m *Material) {
		m.shininess = value
	}
}

func Reflective(value float64) func(*Material) {
	return func (m *Material) {
		m.reflective = value
	}
}

func Transparency(value float64) func(*Material) {
	return func (m *Material) {
		m.transparency = value
	}
}

func Refractive(value float64) func(*Material) {
	return func (m *Material) {
		m.refractiveIndex = value
	}
}

func NewMaterial(options ...func(*Material)) Material {
	material := Material{
		color:           NewColor(1, 1, 1),
		ambient:         0.1,
		diffuse:         0.9,
		specular:        0.9,
		shininess:       200.0,
		reflective:      0.0,
		transparency:    0.0,
		refractiveIndex: 1.0,
	}

	for _, option := range options {
		option(&material)
	}

	return material
}

func Glass() Material {
	glass := NewMaterial()
	glass.transparency = 1.0
	glass.refractiveIndex = 1.5

	return glass
}

func (m1 Material) Equal(m2 Material) bool {
	return m1.color.Eq(m2.color) &&
		(m1.ambient == m2.ambient) &&
		(m1.diffuse == m2.diffuse) &&
		(m1.specular == m2.specular) &&
		(m1.shininess == m2.shininess)
}

// func (m *Material) SetPattern(val *Pattern) {
// 	m.pattern = val
// }

// func (m *Material) SetColor(val Vector) {
// 	m.color = val
// }

// func (m *Material) SetAmbient(val float64) {
// 	m.ambient = val
// }

// func (m *Material) SetDiffuse(val float64) {
// 	m.diffuse = val
// }

// func (m *Material) SetSpecular(val float64) {
// 	m.specular = val
// }

// func (m *Material) SetShininess(val float64) {
// 	m.shininess = val
// }

// func (m *Material) SetReflective(val float64) {
// 	m.reflective = val
// }

// func (m *Material) SetTransparency(val float64) {
// 	m.transparency = val
// }

// func (m *Material) SetRefractiveIndex(val float64) {
// 	m.refractiveIndex = val
// }
