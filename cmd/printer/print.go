package printer

import "github.com/fatih/color"

type Printer interface {
	Info(message string)
}

type printer struct{}

func NewPrinter() Printer {
	return &printer{}
}

func (p *printer) Info(message string) {
	color.Cyan(message)
}
