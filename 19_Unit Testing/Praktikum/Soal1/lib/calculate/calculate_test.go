package calculate

import "testing"

func TestAddition(t *testing.T) {
	if Addition(50, 40) != 90 {
		t.Error("Expected 50 (+) 40 to equal 90")
	}
}

func TestSubstract(t *testing.T) {
	if Substract(30, 10) != 20 {
		t.Error("Expected 30 (-) 10 to equal 20")
	}
}

func TestMult(t *testing.T) {
	if Mult(40, 5) != 200 {
		t.Error("Expected 40 (*) 5 to equal 200")
	}
}

func TestDiv(t *testing.T) {
	if Div(10, 2) != 5 {
		t.Error("Expected 10 (/) 2 to equal 5")
	}
}