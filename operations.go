package main

import "math"

type operator int

const (
	PLUS operator = iota
	MINUS
	MULT
	DIV
	MOD
	POWER
)

var opmap map[string]operator = map[string]operator{
	"+":  PLUS,
	"-":  MINUS,
	"*":  MULT,
	"/":  DIV,
	"%":  MOD,
	"^":  POWER,
	"**": POWER,
}

func readOperator(input string) operator {
	op, ok := opmap[input]
	if ok {
		return op
	} else {
		return -1
	}
}

func (op operator) show() string {
	for key, value := range opmap {
		if value == op {
			return key
		}
	}
	return "UNKNOWN OPERATOR"
}

func (o operator) apply(a float64, b float64) float64 {
	switch o {
	case PLUS:
		return a + b
	case MINUS:
		return a - b
	case MULT:
		return a * b
	case DIV:
		return a / b
	case MOD:
		return math.Mod(a, b)
	case POWER:
		return math.Pow(a, b)
	}
	return a
}

type function int

const (
	LOG function = iota
	LN
	LB
	INVERT
)

var funcmap map[string]function = map[string]function{
	"log":   LOG,
	"log10": LOG,
	"ln":    LN,
	"loge":  LN,
	"lb":    LB,
	"log2":  LB,
	"-":     INVERT,
}

func (f *function) apply(a float64) float64 {
	switch *f {
	case LOG:
		return math.Log10(a)
	case LN:
		return math.Log(a)
	case LB:
		return math.Log2(a)
	case INVERT:
		return -a
	}
	return a
}
