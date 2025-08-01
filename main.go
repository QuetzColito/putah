package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"text/scanner"
	"unicode"
)

type computer interface {
	compute() float64
}

type operation struct {
	left  *computer
	op    operator
	right *computer
}

func (o *operation) compute() float64 {
	x := (*o.left).compute()
	y := (*o.right).compute()
	switch o.op {
	case PLUS:
		return x + y
	case MINUS:
		return x - y
	}
	return -1
}

type application struct {
	argument *computer
	op       function
}

func (a *application) compute() float64 {
	x := (*a.argument).compute()
	switch a.op {
	case LOG:
		return math.Log10(x)
	}
	return -1
}

type literal struct {
	value float64
}

func (l *literal) compute() float64 {
	return l.value
}

type operator int

const (
	PLUS operator = iota
	MINUS
)

type function int

const (
	LOG function = iota
)

func main() {
	args := strings.Join(os.Args[1:], "")
	validate(args)
	var s scanner.Scanner
	s.Init(strings.NewReader(args))
	sum := 0.0
	lastop := PLUS
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		op, isOp := toOperator(s.TokenText())
		switch {
		case isNumber(s.TokenText()):
			fmt.Printf("Found number: %s\n", s.TokenText())
			v, _ := strconv.ParseFloat(s.TokenText(), 64)
			sum = compute(sum, lastop, v)
		case isOp:
			fmt.Printf("Found op: %s\n", s.TokenText())
			lastop = op
		}
	}
	fmt.Printf("sum: %f\n", sum)
}

func compute(a float64, op operator, b float64) float64 {
	switch op {
	case MINUS:
		return a - b
	case PLUS:
		return a + b
	}
	return -1
}

func toOperator(token string) (operator, bool) {
	switch token {
	case "-":
		return MINUS, true
	case "+":
		return PLUS, true
	default:
		return -1, false
	}

}

func validate(str string) bool {
	runes := []rune(str)
	count := 0
	for i, r := range runes {
		switch r {
		case '(':
			count++
		case ')':
			count--
			if count < 0 {
				fmt.Printf("Unexpected ')' at %d", i)
				return false
			}
		}
	}

	if count != 0 {
		fmt.Println("Not All Parens closed")
		return false
	} else {
		return true
	}
}

func isNumber(token string) bool {
	runes := []rune(token)
	return every(runes, func(r rune, _ int, _ []rune) bool { return unicode.IsNumber(r) || r == '.' })
}

func every[T any](slice []T, predicate func(value T, index int, slice []T) bool) bool {
	for i, el := range slice {
		if ok := predicate(el, i, slice); !ok {
			return false
		}
	}
	return true
}
