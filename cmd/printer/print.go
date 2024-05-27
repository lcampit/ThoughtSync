package printer

import (
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type Printer interface {
	Info(message string)
	ConfigMainKey(message string)
	ConfigSubKey(message string)
	CustomError(message string)
	PlainError(err error)
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
	splitKeyValue := strings.Split(message, ":")
	key := splitKeyValue[0]
	value := strings.TrimSpace(splitKeyValue[1])
	color.New(color.FgWhite).Printf("%s: ", key)
	p.PrintConfigValueByType(value)
}

func (p *printer) PrintConfigValueByType(value string) {
	if valueAsBoolean, err := strconv.ParseBool(value); err == nil {
		if valueAsBoolean {
			color.Green(value)
		} else {
			color.Red(value)
		}
	} else {
		color.New(color.FgCyan).Add(color.Bold).Println(value)
	}
}

func (p *printer) CustomError(message string) {
	color.Red(message)
}

func (p *printer) PlainError(err error) {
	color.Red(err.Error())
}
