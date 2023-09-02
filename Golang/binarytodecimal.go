package golang

import (
	"fmt"
	"strconv"
)

func decimalToBinary() {
	var decimal int
	fmt.Print("Enter a decimal number: ") // No new line
	fmt.Scan(&decimal)                    // Takes input and puts it into the variable
	binary := strconv.FormatInt(int64(decimal), 2)
	// Conversion taken from https://www.techieindoor.com/go-program-to-convert-decimal-number-to-binary-in-go/
	// FromatInt takes only 64 bit integers and require conversion
	fmt.Printf("Converted Binary Value: %s\n", binary) // F string

}

func binaryToDecimal() {
	// Conversion taken from https://stackoverflow.com/questions/9271469/go-convert-string-which-represent-binary-number-into-int
	var binary string // based on the example, we take a string that represents the Binary
	fmt.Print("Enter a binary number: ")
	fmt.Scan(&binary)
	conversion, err := strconv.ParseInt(binary, 2, 64) // function returns both err and converted item
	if err != nil {                                    // nil is default response if successful
		fmt.Println("Please Enter a Valid Binary")
		return
	}
	fmt.Printf("Converted Decimal Value: %d\n", conversion)
}
