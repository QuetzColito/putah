package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/scanner"
	"unicode"
)

func main() {
	args := strings.Join(os.Args[1:], "")
	validate(args)
	var s scanner.Scanner
	s.Init(strings.NewReader(args))

	var root tree
	var next *operation
	expectExpression := true
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		switch {
		case root == nil:
			v, _ := strconv.ParseFloat(s.TokenText(), 64)
			root = literal{v}
		case next != nil:
			v, _ := strconv.ParseFloat(s.TokenText(), 64)
			if next.op == MINUS {
				next.right = literal{-v}
				next.op = PLUS
			} else {
				next.right = literal{v}
			}
			root = root.attach(*next)
			next = nil
		case next == nil:
			op := readOperator(s.TokenText())
			next = &operation{op: op}
		}
		expectExpression = !expectExpression
	}
	root.print(0)
	fmt.Printf("sum: %f\n", root.compute())
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
