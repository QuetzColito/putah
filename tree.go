package main

import (
	"fmt"
	"strings"
)

type tree interface {
	compute() float64
	attach(operation) tree
	print(int)
}

type operation struct {
	left  tree
	op    operator
	right tree
}

func (o operation) compute() float64 {
	x := (o.left).compute()
	y := (o.right).compute()
	return o.op.apply(x, y)
}

func (o operation) attach(t operation) tree {
	if t.op < o.op {
		t.left = o
		return t
	} else {
		newRight := o.right.attach(t)
		o.right = newRight
		return o
	}
}

func (o operation) print(offset int) {
	indent := strings.Repeat(">", offset)
	fmt.Printf("%s Op: %s\n", indent, o.op.show())
	o.left.print(offset + 1)
	o.right.print(offset + 1)
}

type application struct {
	argument *tree
	op       function
}

func (a application) attach(t operation) tree {
	t.left = a
	return t
}

func (a application) compute() float64 {
	x := (*a.argument).compute()
	return a.op.apply(x)
}

func (a application) print(offset int) {
}

type literal struct {
	value float64
}

func (l literal) attach(t operation) tree {
	t.left = l
	return t
}

func (l literal) print(offset int) {
	indent := strings.Repeat(">", offset)
	fmt.Printf("%s lit: %f\n", indent, l.value)
}

func (l literal) compute() float64 {
	return l.value
}
