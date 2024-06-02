package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

// func TestArea(t *testing.T) {
// 	rect := Rectangle{12, 6}
// 	got := Area(rect)
// 	want := 72.0
//
// 	if got != want {
// 		t.Errorf("got %.2f, want %.2f", got, want)
// 	}
// }

func TestArea(t *testing.T) {
	shapeTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{
			name:  "Rectangle",
			shape: Rectangle{10, 4},
			want:  40,
		},
		{
			name:  "Circle",
			shape: Circle{10},
			want:  314.1592653589793,
		},
		{
			name:  "Square",
			shape: Square{10, 5},
			want:  51,
		},
	}

	for _, tt := range shapeTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("Area(), got %.2f -> want %.2f", got, tt.want)
			}
		})
	}
}
