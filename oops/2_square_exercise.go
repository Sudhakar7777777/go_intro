package main

import (
	"fmt"
	"log"
)

type Point struct {
	X int
	Y int
}

func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

type Square struct {
	Center Point
	Length int
}

func (s *Square) Move(dx int, dy int) {
	s.Center.Move(dx, dy)
}

func (s *Square) Area() int {
	return s.Length * s.Length
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("Length must be > 0")
	}
	s := &Square{
		Center: Point{x, y},
		Length: length,
	}
	return s, nil
}

func main() {
	s, err := NewSquare(2, 3, 20)
	if err != nil {
		// fmt.Printf("Error: %s\n", err)
		// os.Exit(1)
		log.Fatalf("Error: Can't create square")
	}
	fmt.Println("Square:", s)

	fmt.Println("Area:", s.Area())

	s.Move(2, 2)
	fmt.Println("Moved Square:", s)
}
