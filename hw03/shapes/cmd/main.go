package main

import (
	"fmt"

	"github.com/ptsypyshev/gb-golang-level1-new/hw03/shapes/internal/models"
)

func main() {
	r := models.NewRectangle(3, 4)
	c := models.NewCircle(5)
	fmt.Printf("Rectangle area: %.02f, circle area: %.02f\n", r.Area(), c.Area())
}