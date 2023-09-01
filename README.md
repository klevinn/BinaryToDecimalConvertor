# Binary To Decimal

This is a simple command line interface (CLI) program that can convert decimal numbers to binary and vice versa. There is a both a Python and Golang version in this repository

## Code Explanation

The `decimalToBinary()` function takes a decimal number as input and converts it to binary. The function first converts the decimal number to a string. 

In Python, it uses the `bin()` function to convert the decimal to a binary representation.

In Golang, it uses the `strconv.FormatInt()` function to convert the decimal to a binary representation. 

The binary representation is then printed to the console.

The `binaryToDecimal()` function takes a binary number as input and converts it to decimal. 

In Python, it uses the `int()` function to convert the binary representation to a decimal number

In Golang, it uses the `strconv.ParseInt()` function to binary representation to a decimal number

The decimal number is then printed to the console.

## Packages Used
[fmt](https://pkg.go.dev/fmt)

[strconv](https://pkg.go.dev/strconv)
