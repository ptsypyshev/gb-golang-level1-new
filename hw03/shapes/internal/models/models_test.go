package models

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShape_Area(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		s    *Shape
		want float64
	}{
		{
			name: "Init with {}",
			s:    &Shape{},
			want: 0,
		},
		{
			name: "Init with new",
			s:    new(Shape),
			want: 0,
		},
		{
			name: "Nil Shape",
			s:    nil,
			want: 0,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.s.Area())
		})
	}
}

func TestRectangle_Area(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		r         *Rectangle
		want      float64
		wantPanic bool
	}{
		{
			name: "Init with {}",
			r:    &Rectangle{},
			want: 0,
		},
		{
			name: "Init with new",
			r:    new(Rectangle),
			want: 0,
		},
		{
			name:      "Nil Rectangle",
			r:         nil,
			want:      0,
			wantPanic: true,
		},
		{
			name: "Int Rectangle",
			r:    &Rectangle{sideA: 5, sideB: 3},
			want: 15,
		},
		{
			name: "Float Rectangle",
			r:    &Rectangle{sideA: 5.5, sideB: 3.3},
			want: 18.15,
		},		
		{
			name: "Constructor Rectangle",
			r:    NewRectangle(3, 4.1),
			want: 12.3,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.wantPanic {
				assert.PanicsWithError(t, "runtime error: invalid memory address or nil pointer dereference", func() { tt.r.Area() })
				return
			}
			assert.InDelta(t, tt.want, tt.r.Area(), tt.want*0.001)
		})
	}
}

func TestCircle_Area(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		c         *Circle
		want      float64
		wantPanic bool
	}{
		{
			name: "Init with {}",
			c:    &Circle{},
			want: 0,
		},
		{
			name: "Init with new",
			c:    new(Circle),
			want: 0,
		},
		{
			name:      "Nil Circle",
			c:         nil,
			want:      0,
			wantPanic: true,
		},
		{
			name: "Int Circle",
			c:    &Circle{radius: 5},
			want: 157,
		},
		{
			name: "Float Circle",
			c:    &Circle{radius: 500.5},
			want: 1573141,
		},		
		{
			name: "Constructor Circle",
			c:    NewCircle(4.1),
			want: 105.6,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.wantPanic {
				assert.PanicsWithError(t, "runtime error: invalid memory address or nil pointer dereference", func() { tt.c.Area() })
				return
			}
			assert.InDelta(t, tt.want, tt.c.Area(), tt.want*(math.Pi-3.14))
		})
	}
}
