package main

type Circle struct {
	R            float32
	BGColour     string
	Border       float32
	BorderColour string
}

func NewCircle() *Circle {
	return &Circle{}
}

func (c *Circle) SetR(r float32) IShape {
	c.R = r
	return c
}

func (c *Circle) SetBGColour(colour string) IShape {
	c.BGColour = colour
	return c
}

func (c *Circle) SetBorder(b float32) IShape {
	c.Border = b
	return c
}

func (c *Circle) SetBorderColour(colour string) IShape {
	c.BorderColour = colour
	return c
}

type IShape interface {
	SetR(r float32) IShape
	SetBGColour(colour string) IShape
	SetBorder(b float32) IShape
	SetBorderColour(colour string) IShape
}

// builder pattern
// Create Circle
// GBColour ->Red
// Border = 10
// BrderColour = Blue
