package samples

import (
	"fmt"
	"strings"
	"unicode"
)

type UserInput interface {
	Add(rune)
	GetValue() string
}

type NumericInput struct {
	input string
}

func (ni *NumericInput) Add(r rune) {
	if !unicode.IsDigit(r) {
		return
	}
	ni.input = strings.Join([]string{ni.input, string(r)}, "")
	fmt.Println("Input changed to ", ni.input)
}

func (ni *NumericInput) GetValue() string {
	return ni.input
}
