package Golang

import (
	"fmt"
	"time"
)

func bruteforce(pw string) (int, string) {
	// attempts := 0
	// chars := strings.Split("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~", "")
	// for i := 1; i <= len(pw); i++ {
	// 	for _, guess := range cartesian.Iter(chars, i) {
	// 		attempts++
	// 		if strings.Join(guess, "") == pw {
	// 			return attempts, strings.Join(guess, "")
	// 		}
	// 	}
	// }
	// return attempts, ""
}

func main() {
	startTime := time.Now()
	var pw string
	fmt.Print("Input the password to crack: ")
	fmt.Scanln(&pw)
	attempts, guess := bruteforce(pw)
	if guess != "" {
		fmt.Printf("Password cracked in %d attempts. The password is %s.\n", attempts, guess)
	} else {
		fmt.Printf("Password not cracked after %d attempts.\n", attempts)
	}
	fmt.Printf("Time taken to complete execution: %s\n", time.Since(startTime))
}
