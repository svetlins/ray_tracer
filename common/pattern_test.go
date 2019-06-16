package common

import (
	"testing"
)

func TestStripeObjectTransform(t *testing.T) {
	stripe := NewPattern()
	sphere := NewSphere()
	sphere.SetTransform(Scaling(2, 2, 2))

	if stripe.AtObject(sphere, NewPoint(1.5, 0, 0)).NotEq(White()) {
		t.Errorf("Stripe At not working")
	}
}

func TestStripePatternTransform(t *testing.T) {
	stripe := NewPattern()
	sphere := NewSphere()
	stripe.SetTransform(Scaling(2, 2, 2))

	if stripe.AtObject(sphere, NewPoint(1.5, 0, 0)).NotEq(White()) {
		t.Errorf("Stripe At not working")
	}
}

func TestStripePatternAndObjectTransform(t *testing.T) {
	stripe := NewPattern()
	stripe.SetTransform(Translation(0.5, 0, 0))

	sphere := NewSphere()
	sphere.SetTransform(Scaling(2, 2, 2))

	if stripe.AtObject(sphere, NewPoint(1.5, 0, 0)).NotEq(White()) {
		t.Errorf("Stripe At not working")
	}
}

// func TestGradient(t *testing.T) {
// 	black := Black()
// 	white := White()
// 	gray := black.Add(white).Scale(0.5)

// 	stripe := NewPattern()
// 	stripe.SetValueFunc(Gradient())

// 	if stripe.AtObject(NewSphere(), NewPoint(1.5, 0, 0)).NotEq(white) {
// 		t.Errorf("Stripe At not working")
// 	}

// 	if stripe.AtObject(NewSphere(), NewPoint(1, 0, 0)).NotEq(black) {
// 		t.Errorf("Stripe At not working")
// 	}

// 	if stripe.AtObject(NewSphere(), NewPoint(1.25, 0, 0)).NotEq(gray) {
// 		t.Errorf("Stripe At not working")
// 	}
// }

// func TestRing(t *testing.T) {
// 	black := Black()
// 	white := White()

// 	stripe := NewPattern()
// 	stripe.SetValueFunc(Ring())

// 	if stripe.AtObject(NewSphere(), NewPoint(0, 0, 0)).NotEq(white) {
// 		t.Errorf("Ring At not working")
// 	}

// 	if stripe.AtObject(NewSphere(), NewPoint(1, 0, 0)).NotEq(black) {
// 		t.Errorf("Ring At not working")
// 	}

// 	if stripe.AtObject(NewSphere(), NewPoint(0, 0, 1)).NotEq(black) {
// 		t.Errorf("Ring At not working")
// 	}

// 	if stripe.AtObject(NewSphere(), NewPoint(0.78, 0, 0.78)).NotEq(black) {
// 		t.Errorf("Ring At not working")
// 	}
// }

// func TestCheckers1(t *testing.T) {
// 	black := Black()
// 	white := White()

// 	stripe := NewPattern()
// 	stripe.SetValueFunc(Checkers())

// 	if stripe.AtObject(NewSphere(), NewPoint(0, 0, 0)).NotEq(white) {
// 		t.Errorf("Ring At not working")
// 	}

// 	if stripe.AtObject(NewSphere(), NewPoint(0.99, 0, 0)).NotEq(white) {
// 		t.Errorf("Ring At not working")
// 	}

// 	if stripe.AtObject(NewSphere(), NewPoint(1.01, 0, 0)).NotEq(black) {
// 		t.Errorf("Ring At not working")
// 	}
// }

// func TestCheckers2(t *testing.T) {
// 	black := Black()
// 	white := White()

// 	stripe := NewPattern()
// 	stripe.SetValueFunc(Checkers())

// 	if stripe.AtObject(NewSphere(), NewPoint(0, 0, 0)).NotEq(white) {
// 		t.Errorf("Ring At not working")
// 	}

// 	if stripe.AtObject(NewSphere(), NewPoint(0, 0.99, 0)).NotEq(white) {
// 		t.Errorf("Ring At not working")
// 	}

// 	if stripe.AtObject(NewSphere(), NewPoint(0, 1.01, 0)).NotEq(black) {
// 		t.Errorf("Ring At not working")
// 	}
// }

// func TestCheckers3(t *testing.T) {
// 	black := Black()
// 	white := White()

// 	stripe := NewPattern()
// 	stripe.SetValueFunc(Checkers())

// 	if stripe.AtObject(NewSphere(), NewPoint(0, 0, 0)).NotEq(white) {
// 		t.Errorf("Ring At not working")
// 	}

// 	if stripe.AtObject(NewSphere(), NewPoint(0, 0, 0.99)).NotEq(white) {
// 		t.Errorf("Ring At not working")
// 	}

// 	if stripe.AtObject(NewSphere(), NewPoint(0, 0, 1.01)).NotEq(black) {
// 		t.Errorf("Ring At not working")
// 	}
// }
