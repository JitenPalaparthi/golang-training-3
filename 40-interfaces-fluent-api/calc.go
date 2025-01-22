package main

type Calc struct {
	data int
}

func New(d int) *Calc {
	return &Calc{data: d}
}

func (c *Calc) Add(n int) ICalc {
	c.data += n
	return c
}

func (c *Calc) Sub(n int) ICalc {
	c.data -= n
	return c
}

func (c *Calc) Mul(n int) ICalc {
	c.data *= n
	return c
}

func (c *Calc) Div(n int) ICalc {
	c.data /= n
	return c
}

func (c *Calc) Get() int {
	return c.data
}
