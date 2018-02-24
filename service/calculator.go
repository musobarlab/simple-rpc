package service

import (
	"errors"
)

//Inputs contains X and Y
type Inputs struct {
	X, Y float64
}

//Result
type Result struct {
	Z     float64
	Error error
}

//calculator struct, implement all service method
type calculator struct {
	Name string
}

//NewCalculator function, calculator's constructor
func NewCalculator(name string) *calculator {
	return &calculator{name}
}

//Add operation
func (c *calculator) Add(i *Inputs, r *Result) error {
	r.Z = i.X + i.Y
	return nil
}

//Div operation
func (c *calculator) Div(i *Inputs, r *Result) error {
	if i.Y == 0 {
		return errors.New("Error, Divide by Zero")
	}
	r.Z = i.X / i.Y
	return nil
}

//Mul operation
func (c *calculator) Mul(i *Inputs, r *Result) error {
	r.Z = i.X * i.Y
	return nil
}

//Sub operation
func (c *calculator) Sub(i *Inputs, r *Result) error {
	r.Z = i.X - i.Y
	return nil
}
