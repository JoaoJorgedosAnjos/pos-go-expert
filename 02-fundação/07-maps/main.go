package main

import "fmt" // Import the fmt package for formatted I/O (like printing to the console)

func main() {
	// --- Declaring and Initializing Maps ---

	// A map is a powerful data structure in Go that stores key-value pairs.
	// Here, we're creating a map named 'salary' where:
	// - The keys are of type 'string' (representing employee names).
	// - The values are of type 'int' (representing their salaries).
	// We're also initializing it with some initial data.
	salary := map[string]int{
		"Maria": 1600,
		"Isis":  3000,
		"João":  1500, // Initial entry for João
	}

	// You can uncomment the lines below to see the map's contents and a specific value.
	// fmt.Println("Initial salaries:", salary)
	// fmt.Println("João's initial salary:", salary["João"])

	// --- Deleting an Entry from a Map ---

	// The 'delete' function is used to remove a key-value pair from a map.
	// Here, we're removing "João" from the 'salary' map.
	delete(salary, "João")
	fmt.Println("\nAfter deleting João, salaries:", salary)

	// After deletion, trying to access a deleted key will return the zero value
	// for the value type (0 for int in this case), and 'false' for the optional second return value.
	// fmt.Println("João's salary after deletion (will be 0):", salary["João"])

	// --- Adding a New Entry to a Map ---

	// You can add a new key-value pair to a map simply by assigning a value
	// to a new key. If the key already exists, its value will be updated.
	salary["Lucios"] = 6000 // Adding Lucios to the map
	fmt.Println("After adding Lucios, salaries:", salary)

	// --- Different Ways to Declare Empty Maps ---

	// 'make' is often used to create maps (and slices, channels).
	// When using 'make', you can optionally specify an initial capacity for performance
	// if you know roughly how many elements the map will hold.
	sal := make(map[string]int)
	sal["Teste"] = 20 // Adding an entry to 'sal'
	fmt.Println("\n'sal' map (created with make):", sal)

	// This is a shorthand for declaring an empty map. It's concise and commonly used
	// when you don't need to specify an initial capacity.
	sal1 := map[string]int{}
	sal1["Teste 2"] = 25 // Adding an entry to 'sal1'
	fmt.Println("'sal1' map (created with literal):", sal1)

	// --- Iterating Over a Map with a 'for...range' Loop ---

	fmt.Println("\n--- Iterating through salaries (name and salary) ---")
	// The 'for...range' loop is the idiomatic way to iterate over maps in Go.
	// It returns two values for each iteration:
	// 1. The key (here, 'name').
	// 2. The value (here, 'salary').
	// The order of iteration over a map is not guaranteed to be the same each time.
	for name, currentSalary := range salary { // Renamed 'salary' to 'currentSalary' to avoid shadowing the map variable
		fmt.Printf("O salário do %s é R$%d\n", name, currentSalary)
	}

	// --- Iterating Over a Map (Values Only) using the Blank Identifier ---

	fmt.Println("\n--- Iterating through salaries (salary only) ---")
	// If you only need the values from a map and don't care about the keys,
	// you can use the **blank identifier** (`_`) in place of the key variable.
	// This tells Go that you intend to ignore that value.
	for _, currentSalary := range salary { // Again, using 'currentSalary' for clarity
		fmt.Printf("O salário é R$%d\n", currentSalary)
	}
}