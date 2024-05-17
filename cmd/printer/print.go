package printer

import "github.com/fatih/color"

type Printer interface {
	Info(message string)
	ConfigMainKey(message string)
	ConfigSubKey(message string)
}

type printer struct{}

func NewPrinter() Printer {
	return &printer{}
}

func (p *printer) Info(message string) {
	color.Cyan(message)
}

func (p *printer) ConfigMainKey(message string) {
	color.New(color.FgWhite).Add(color.Bold).Add(color.BgBlack).Println(message)
}

func (p *printer) ConfigSubKey(message string) {
	color.New(color.FgCyan).Add(color.Bold).Println(message)
}
