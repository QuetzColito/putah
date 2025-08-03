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

func main() {
	args := strings.Join(os.Args[1:], "") + ")"
	validate("(" + args)
	var s scanner.Scanner
	s.Init(strings.NewReader(args))

	result, _ := parseParens(s)
	fmt.Printf("result: %f\n", result)
}

func parseParens(s scanner.Scanner) (float64, scanner.Scanner) {
	var root tree
	var next *operation
	var functions []function
	for tok := s.Scan(); tok != ')'; tok = s.Scan() {
		token := s.TokenText()
		fmt.Println(token)
		if next != nil || root == nil {
			// Read Expression
			var value tree // stores the result of the Expression
			switch {
			case readOperator(token) > 0:
				return math.NaN(), s
			case token == "(":
				var v float64
				v, s = parseParens(s) // recurse
				value = literal{v}
			case !isNumber(token):
				// Read Function
				functions = append(functions, readFunction(token))
				continue
			default:
				// Read Number
				v, err := strconv.ParseFloat(token, 64)
				if err != nil {
					value = literal{0}
				} else {
					value = literal{v}
				}
			}

			for i := len(functions) - 1; i >= 0; i-- {
				value = application{op: functions[i], argument: value}
			}
			functions = nil

			if root == nil {
				root = value
			} else {
				if next.op == MINUS {
					next.right = application{op: INVERT, argument: value}
					next.op = PLUS
				} else {
					next.right = value
				}
				root = root.attach(*next)
				next = nil
			}
		} else {
			// Read Operator
			op := readOperator(token)
			if op < 0 {
				return math.NaN(), s
			}
			next = &operation{op: op}
		}
	}
	if root == nil {
		return math.NaN(), s
	} else {
		return root.compute(), s
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
