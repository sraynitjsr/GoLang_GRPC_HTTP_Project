package calcapi

import (
	"context"
	"log"

	calc "github.com/sraynitjsr/gen/calc"
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
	return
}

// Sub implements sub.
func (s *calcsrvc) Sub(ctx context.Context, p *calc.SubPayload) (res int, err error) {
	s.logger.Print("calc.sub")
	return
}

// Multiply implements multiply.
func (s *calcsrvc) Multiply(ctx context.Context, p *calc.MultiplyPayload) (res int, err error) {
	s.logger.Print("calc.multiply")
	return
}

// Divide implements divide.
func (s *calcsrvc) Divide(ctx context.Context, p *calc.DividePayload) (res int, err error) {
	s.logger.Print("calc.divide")
	return
}
