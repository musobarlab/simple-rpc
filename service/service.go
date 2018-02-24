package service

const CalculatorName = "Calculator"

//Service interface
type Service interface {
	Add(*Inputs, *Result) error
	Div(*Inputs, *Result) error
	Mul(*Inputs, *Result) error
	Sub(*Inputs, *Result) error
}
