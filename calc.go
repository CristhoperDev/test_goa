package calcapi

import (
	calc "calc/gen/calc"
	"context"
	"log"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calc.Service {
	return &calcsrvc{logger}
}

// Add implements add.
func (s *calcsrvc) Add(ctx context.Context, p *calc.AddPayload) (res int, err error) {
	s.logger.Print("calc.add")
	return p.A + p.B, nil
}

// Minus implements minus.
func (s *calcsrvc) Minus(ctx context.Context, p *calc.MinusPayload) (res int, err error) {
	s.logger.Print("calc.minus")
	return p.A - p.B, nil
}

// Multiplication implements multiplication.
func (s *calcsrvc) Multiplication(ctx context.Context, p *calc.MultiplicationPayload) (res int, err error) {
	s.logger.Print("calc.multiplication")
	return p.A * p.B, nil
}
