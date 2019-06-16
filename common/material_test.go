package common

import (
	"testing"
)

func TestNewMaterial(t *testing.T) {
	m := NewMaterial()

	if !m.color.Eq(NewColor(1, 1, 1)) {
		t.Errorf("New material color not working")
	}

	if m.ambient != 0.1 {
		t.Errorf("New material color not working")
	}

	if m.diffuse != 0.9 {
		t.Errorf("New material color not working")
	}

	if m.specular != 0.9 {
		t.Errorf("New material color not working")
	}

	if m.shininess != 200.0 {
		t.Errorf("New material color not working")
	}

	if m.reflective != 0.0 {
		t.Errorf("New material color not working")
	}

	if m.transparency != 0.0 {
		t.Errorf("New material color not working")
	}

	if m.refractiveIndex != 1.0 {
		t.Errorf("New material color not working")
	}
}
