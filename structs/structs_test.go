package structs

import "testing"

// Requirement:
// Calculate permimeter of a rectangle

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{4.0, 5.0}
	got := Perimeter(rectangle)
	want := 18.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}

}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 4.0, Height: 5.0}, hasArea: 20},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 40, Height: 10}, hasArea: 200},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.hasArea {
				t.Errorf("got %.2f has area %.2f", got, test.hasArea)
			}

		})
	}

}
