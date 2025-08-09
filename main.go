package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/scanner"
	"unicode"
)

func main() {
	args := strings.ToLower(strings.Join(os.Args[1:], "") + ")")
	args = strings.ReplaceAll(args, ",", ".") // allow for commas :P
	args = strings.ReplaceAll(args, "x", "*") // Only works because none of the functions contain x

	// Sanitize so go doesnt think 2pi is scientific notation >.>
	re := regexp.MustCompile(`(\d)([pe])`)
	args = re.ReplaceAllString(args, `$1 $2`)

	// make sure parens match
	valid := validate("(" + args)
	if valid {
		var s scanner.Scanner
		s.Init(strings.NewReader(args))
		result, _ := parseParens(s)
		fmt.Println(result)
	} else {
		fmt.Println(math.NaN())
	}
}

func parseParens(s scanner.Scanner) (float64, scanner.Scanner) {
	var root tree
	var next *operation
	var functions []function

	// The Loop
	for tok := s.Scan(); tok != ')'; tok = s.Scan() {
		token := s.TokenText()
	readToken:
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
				// Try to Read as Constant
				constant, ok := constants[token]
				if ok {
					value = literal{constant}
				} else {
					// Read Function, default is identity
					functions = append(functions, readFunction(token))
					continue
				}
			default:
				// Read Number
				v, err := strconv.ParseFloat(token, 64)
				if err != nil {
					value = literal{0}
				} else {
					value = literal{v}
				}
			}

			// If we got here we got a value, so apply all functions we found along the way
			for i := len(functions) - 1; i >= 0; i-- {
				value = application{op: functions[i], argument: value}
			}
			functions = nil

			// Attach Operation to the tree
			if root == nil {
				root = value
			} else {
				// No idea why i cant do minus normally
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
				// If not an Operator, read as an 'invisible mult'
				next = &operation{op: MULT}
				goto readToken
			}
			next = &operation{op: op}
		}
	}
	// No Input -> NaN
	if root == nil {
		return math.NaN(), s
	} else {
		return root.compute(), s
	}
}

func isNumber(token string) bool {
	runes := []rune(token)
	for _, rune := range runes {
		if !unicode.IsNumber(rune) && !(rune == '.') {
			return false
		}
	}
	return true
}

func validate(str string) bool {
	runes := []rune(str)
	count := 0
	for _, r := range runes {
		switch r {
		case '(':
			count++
		case ')':
			count--
			if count < 0 {
				return false
			}
		}
	}

	if count != 0 {
		return false
	} else {
		return true
	}
}
