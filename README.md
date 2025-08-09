# Putah - Just a Calculator

Might as well add a Readme.

When called, concatenates all arguments and tries to evaluate them with some basic math.
Made it to be used in my Quickshell Widgets.

## Behaviour:

- Support for basic [Operators](#Supported-Operators), some 1 Argument [Functions](#Supported-Functions) and [Constants](#Supported-Constants).
- The functions behave like their golang math equivalents.
- Support for Parentheses.
- Parentheses around function arguments are optional (so `sqrt 4` is valid and evaluates to 2).
- When an Operator is omitted between 2 distinguishable expressions, Multiplication is assumed (`4(2-3)` = `-4`).
- Uses Floating point math with no special handling so expect some inaccuracies (`sin(pi)` != `0`)
- All Text is Case-insensitive.
- Returns NaN when there is a Syntax Error.
- In theory Operators, Functions and Constants are pretty extensible in `operations.go`, but I cant think of any more to add.

## Supported Operators

`-`, `+`, `*`/`x`, `/`, `%`, `^`/`**`

## Supported Functions

`log`/`log10`, `ln`/`loge`, `lb`/`log2`, `sqrt`, `sin`, `asin`, `cos`, `acos`, `tan`, `atan`, `abs`, `-`

## Supported Constants

`pi`, `e`, `phi`
