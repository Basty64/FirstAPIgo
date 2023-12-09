package main

import "fmt"

type Circle struct {
	x, y, r int
}

func (c *Circle) square() int {
	sq := c.x * c.y * c.r
	fmt.Println(sq)
	return sq
}

func main() {
	c := Circle{1, 2, 3}
	c.square()

	a := Circle{2, 3, 5}
	a.square()
}
