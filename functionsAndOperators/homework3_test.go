package homework3

import (
	"testing"
)

func TestSubtract(t *testing.T) {
	tests := []struct {
		x    int
		y    int
		want int
	}{
		{
			x:    0,
			y:    1,
			want: -1,
		},
		{
			x:    2,
			y:    1,
			want: 1,
		},
		{
			x:    7,
			y:    3,
			want: 4,
		},
	}

	for _, tc := range tests {
		got := subtract(tc.x, tc.y)
		if tc.want != got {
			t.Fatalf("subtract(%d, %d): want %d; got %d", tc.x, tc.y, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		x    int
		y    int
		want int
	}{
		{
			x:    0,
			y:    1,
			want: 0,
		},
		{
			x:    2,
			y:    1,
			want: 2,
		},
		{
			x:    7,
			y:    3,
			want: 21,
		},
		{
			x:    -4,
			y:    6,
			want: -24,
		},
	}

	for _, tc := range tests {
		got := multiply(tc.x, tc.y)
		if tc.want != got {
			t.Fatalf("multiply(%d, %d): want %d; got %d", tc.x, tc.y, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		x    int
		y    int
		want int
	}{
		{
			x:    0,
			y:    1,
			want: 0,
		},
		{
			x:    2,
			y:    1,
			want: 2,
		},
		{
			x:    7,
			y:    3,
			want: 2,
		},
	}

	for _, tc := range tests {
		got := divide(tc.x, tc.y)
		if tc.want != got {
			t.Fatalf("divide(%d, %d): want %d; got %d", tc.x, tc.y, tc.want, got)
		}
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		x    int
		y    int
		z    int
		want int
	}{
		{
			x:    0,
			y:    1,
			z:    2,
			want: 3,
		},
		{
			x:    2,
			y:    1,
			z:    -1,
			want: 2,
		},
		{
			x:    7,
			y:    3,
			z:    3,
			want: 13,
		},
	}

	for _, tc := range tests {
		got := sum(tc.x, tc.y, tc.z)
		if tc.want != got {
			t.Fatalf("sum(%d, %d, %d): want %d; got %d", tc.x, tc.y, tc.z, tc.want, got)
		}
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		x    int
		want bool
	}{
		{
			x:    0,
			want: true,
		},
		{
			x:    1,
			want: false,
		},
		{
			x:    -2,
			want: true,
		},
	}

	for _, tc := range tests {
		got := isEven(tc.x)
		if tc.want != got {
			t.Fatalf("isEven(%d): want %v; got %v", tc.x, tc.want, got)
		}
	}
}

func TestIsOdd(t *testing.T) {
	tests := []struct {
		x    int
		want bool
	}{
		{
			x:    0,
			want: false,
		},
		{
			x:    7,
			want: true,
		},
		{
			x:    -2,
			want: false,
		},
	}

	for _, tc := range tests {
		got := isOdd(tc.x)
		if tc.want != got {
			t.Fatalf("isOdd(%d): want %v; got %v", tc.x, tc.want, got)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		x    int
		y    int
		want int
	}{
		{
			x:    -5,
			y:    1,
			want: 1,
		},
		{
			x:    2,
			y:    1,
			want: 2,
		},
		{
			x:    7,
			y:    7,
			want: 7,
		},
	}

	for _, tc := range tests {
		got := max(tc.x, tc.y)
		if tc.want != got {
			t.Fatalf("max(%d, %d): want %d; got %d", tc.x, tc.y, tc.want, got)
		}
	}
}

func TestAnd(t *testing.T) {
	tests := []struct {
		x    bool
		y    bool
		want bool
	}{
		{
			x:    true,
			y:    true,
			want: true,
		},
		{
			x:    true,
			y:    false,
			want: false,
		},
		{
			x:    false,
			y:    true,
			want: false,
		},
		{
			x:    false,
			y:    false,
			want: false,
		},
	}

	for _, tc := range tests {
		got := and(tc.x, tc.y)
		if tc.want != got {
			t.Fatalf("and(%v, %v): want %v; got %v", tc.x, tc.y, tc.want, got)
		}
	}
}

func TestOr(t *testing.T) {
	tests := []struct {
		x    bool
		y    bool
		want bool
	}{
		{
			x:    true,
			y:    true,
			want: true,
		},
		{
			x:    true,
			y:    false,
			want: true,
		},
		{
			x:    false,
			y:    true,
			want: true,
		},
		{
			x:    false,
			y:    false,
			want: false,
		},
	}

	for _, tc := range tests {
		got := or(tc.x, tc.y)
		if tc.want != got {
			t.Fatalf("or(%v, %v): want %v; got %v", tc.x, tc.y, tc.want, got)
		}
	}
}

func TestNot(t *testing.T) {
	tests := []struct {
		x    bool
		want bool
	}{
		{
			x:    true,
			want: false,
		},
		{
			x:    false,
			want: true,
		},
	}

	for _, tc := range tests {
		got := not(tc.x)
		if tc.want != got {
			t.Fatalf("not(%v): want %v; got %v", tc.x, tc.want, got)
		}
	}
}
