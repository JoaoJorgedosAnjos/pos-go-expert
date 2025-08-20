package main

import (
	"errors" // Import the 'errors' package to create custom error messages.
	"fmt"    // Import the 'fmt' package for formatted I/O (like printing to the console).
)

func main() {
	// Call the 'sum' function and assign its two return values to 'valor' and 'err'.
	// In Go, it's a common pattern for functions that might fail to return a second value of type 'error'.
	valor, err := sum(50, 10) // Example 1: sum will be 60, triggering an error.

	// Check if an error occurred.
	// If 'err' is not 'nil' (which means an error was returned), print the error message.
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: A soma é maior que 50
		return                   // It's good practice to exit the function or handle the error appropriately
		// to prevent further execution with potentially invalid data.
	}

	// If no error occurred, print the calculated 'valor'.
	fmt.Println("Result (no error):", valor) // This line will only execute if 'err' is nil.

	// Let's try another call where the sum is less than 50.
	valor2, err2 := sum(10, 20) // Example 2: sum will be 30, no error.
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}
	fmt.Println("Result (no error):", valor2) // Output: Result (no error): 30
}

// sum is a function that takes two integers (a and b) and returns an integer
// and an error. This is a common Go idiom for functions that might fail.
func sum(a, b int) (int, error) {
	// Check if the sum of 'a' and 'b' is greater than or equal to 50.
	if a+b >= 50 {
		// If the condition is met, return 0 (or any appropriate zero value for the primary result)
		// and a new error created using 'errors.New()'.
		return 0, errors.New("A soma é maior que 50")
	}

	// If the condition is not met (i.e., the sum is less than 50),
	// return the actual sum and 'nil' for the error, indicating no error occurred.
	return a + b, nil
}

/*
--- Old Examples (commented out for clarity, but explained below) ---

package main

import "fmt"

func main() {
    fmt.Println(sumOld(1, 2))      // Calls the simple sum function
    fmt.Println(sum1Old(100, 2))   // Calls the sum1Old function with multiple returns
}

// sumOld is a simple function that takes two integers and returns their sum.
// func sumOld(a, b int) int {
//     return a + b
// }

// --- Parameter Declaration Styles ---
// In Go, if consecutive parameters in a function have the same type,
// you only need to specify the type for the last parameter.
// Both of these declarations are valid for 'sumOld':
// func sum(a int, b int) int { ... } // Explicitly stating type for each parameter
// func sum(a, b int) int { ... }     // Shorthand for same-type parameters

// sum1Old is a function that returns multiple values: an integer (the sum)
// and a boolean (indicating if the sum is >= 50).
// func sum1Old(a, b int) (int, bool) {
//     if a+b >= 50 {
//         return a + b, true // Returns the sum and 'true'
//     }
//     return a + b, false    // Returns the sum and 'false'
// }
*/