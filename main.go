package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Shape interface{ Area() float64 }
type Circle struct{ Radius float64 }
type Square struct{ Size float64 }

func (c Circle) Area() float64 { return 3.14 * c.Radius * c.Radius }
func (c Square) Area() float64 { return c.Size * c.Size }

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	kind := sc.Text()
	sc.Scan()
	dim, _ := strconv.ParseFloat(sc.Text(), 64)
	var s Shape

	switch kind {
	case "circle":
		s = Circle{Radius: dim}
	case "square":
		s = Square{Size: dim}
	}

	if s != nil {
		fmt.Printf("%.2f\n", s.Area())
	}
}
