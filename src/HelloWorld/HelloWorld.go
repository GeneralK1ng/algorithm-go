package HelloWorld

// To be executable, the package name must be main
// You can't have two package definitions in the same folder

/*
	1. fmt is the namespace that imports formatting functions, such as Printf() and Sprintf().
	2. runtime is the namespace containing functions that retrieve information about the GO runtime, such as version,
	   CPU, memory and tracing.
	3. The one and only function main() is the entry point of the program.
	4. Once main() is called, the program will exit.
*/
import (
	"fmt"
	"runtime"
	"time"
)

func main() { // It is important to note that '{' cannot be used in a single line
	// Hello World
	// Println() is a function that prints a string followed by a \n.
	fmt.Println("Hello World")

	// Go version
	// runtime.Version() returns the version of the GO runtime
	fmt.Printf("Go version: %s\n", runtime.Version())

	// time
	// time.Now() returns the current time
	fmt.Printf("Time: %s\n", time.Now()) // In a Go program, a line represents the end of a statement

	/*
		Go language variable names are composed of letters, numbers, and underscores,
		and the first character cannot be a number.

		1. `var v_name [type] = [value]`
		2. `var v_name = [value]`
		3. `v_name := [value]`

	*/
	var a = 10
	var b = 20
	c := 30
	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("Sum = %d\n", a+b)
	fmt.Printf("c = %d\n", c)

	/*
		The data types in a constant can only be Boolean, numeric (integer, floating-point, and complex), and string.
		Constant definition format:

		1. `const v_name [type] = [value]`
		2. `const v_name = [value]`
	*/

	const pi float64 = 3.14159265358979323846264338327950288419716
	// The type of radius also needs to be declared as float64 in order to perform correct multiplication with pi.
	radius := float64(10)
	area := pi * radius * radius
	fmt.Printf("Area = %f\n", area)

}
