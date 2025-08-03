package main

import "math"

type operator int

const (
	MINUS operator = iota
	PLUS
	MULT
	DIV
	MOD
	POWER
)

var opmap map[string]operator = map[string]operator{
	"-": MINUS,
	"+": PLUS,
	"*": MULT,
	"/": DIV,
	"%": MOD,
	"^": POWER,
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
	SQRT
	SIN
	ASIN
	COS
	ACOS
	TAN
	ATAN
	ABS
	INVERT
)

var funcmap map[string]function = map[string]function{
	"log":   LOG,
	"log10": LOG,
	"ln":    LN,
	"loge":  LN,
	"lb":    LB,
	"log2":  LB,
	"sqrt":  SQRT,
	"sin":   SIN,
	"asin":  ASIN,
	"cos":   COS,
	"acos":  ASIN,
	"tan":   TAN,
	"atan":  ATAN,
	"abs":   ABS,
	"-":     INVERT,
}

func readFunction(input string) function {
	op, ok := funcmap[input]
	if ok {
		return op
	} else {
		return -1
	}
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
	case SQRT:
		return math.Sqrt(a)
	case SIN:
		return math.Sin(a)
	case ASIN:
		return math.Asin(a)
	case COS:
		return math.Cos(a)
	case ACOS:
		return math.Acos(a)
	case TAN:
		return math.Tan(a)
	case ATAN:
		return math.Atan(a)
	case ABS:
		return math.Asin(a)
	}
	return a
}
