package homework4

import (
	"reflect"
	"testing"
)

func TestZeroThroughTen(t *testing.T) {
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := zeroThroughTen()

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("zeroThroughTen(): want %v; got %v", want, got)
	}
}

func TestZeroThroughTenDecreasing(t *testing.T) {
	want := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	got := zeroThroughTenDecreasing()

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("zeroThroughTenDecreasing(): want %v; got %v", want, got)
	}
}

func TestZeroThroughN(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{
			n:    1,
			want: []int{0},
		},
		{
			n:    0,
			want: []int{},
		},
		{
			n:    5,
			want: []int{0, 1, 2, 3, 4},
		},
	}

	for _, tc := range tests {
		got := zeroThroughN(tc.n)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("zeroThroughN(%d): want %v; got %v", tc.n, tc.want, got)
		}
	}
}

func TestEvensInRange(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want []int
	}{
		{
			a:    0,
			b:    5,
			want: []int{0, 2, 4},
		},
		{
			a:    3,
			b:    3,
			want: []int{},
		},
		{
			a:    6,
			b:    24,
			want: []int{6, 8, 10, 12, 14, 16, 18, 20, 22, 24},
		},
	}

	for _, tc := range tests {
		got := evensInRange(tc.a, tc.b)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("evensInRange(%d, %d): want %v; got %v", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSumThroughTen(t *testing.T) {
	want := 55
	got := sumThroughTen()
	if want != got {
		t.Fatalf("sumThroughTen(): want %d; got %d", want, got)
	}
}

func TestSumThroughN(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{
			n:    10,
			want: 55,
		},
		{
			n:    0,
			want: 0,
		},
		{
			n:    100,
			want: 5050,
		},
		{
			n:    21,
			want: 231,
		},
	}

	for _, tc := range tests {
		got := sumThroughN(tc.n)
		if tc.want != got {
			t.Fatalf("sumThroughN(%d): want %d; got %d", tc.n, tc.want, got)
		}
	}
}

func TestProductThroughTen(t *testing.T) {
	want := 3628800
	got := productThroughTen()

	if want != got {
		t.Fatalf("productThroughTen(): want %d; got %d", want, got)
	}
}

func TestProductThroughN(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{
			n:    10,
			want: 3628800,
		},
		{
			n:    0,
			want: 1,
		},
		{
			n:    5,
			want: 120,
		},
	}

	for _, tc := range tests {
		got := productThroughN(tc.n)
		if tc.want != got {
			t.Fatalf("productThroughN(%d): want %d; got %d", tc.n, tc.want, got)
		}
	}
}

func TestIterationsToN(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{
			n:    10,
			want: 4,
		},
		{
			n:    120,
			want: 5,
		},
		{
			n:    121,
			want: 6,
		},
		{
			n:    3628799,
			want: 10,
		},
	}

	for _, tc := range tests {
		got := iterationsToN(tc.n)
		if tc.want != got {
			t.Fatalf("iterationsToN(%d): want %d; got %d", tc.n, tc.want, got)
		}
	}
}
