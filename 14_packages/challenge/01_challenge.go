package main

import (
	"arithmetic"
	"fmt"
)

func main() {
	/*
			Exercise #1
			Consider this function:

		// Checks if a number is prime or not
		func IsPrime(num int) bool {
		    for i := 2; i < int(num/2); i++ {
		        if num%i == 0 {
		            return false
		        }
		    }
		    return true
		}
			Create a package called arithmetic in $GOPATH/packages/mypackages and paste the above function in a file of package.
			Create an executable Go application, import the package arithmetic and call the function IsPrime.
	*/
	fmt.Println(arithmetic.IsPrime(7))

	/*
			Exercise #2
		Add a new file to the arithmetic package and paste the following code in the file:

		package mymath
		func Factorial(n int) int {
		    var f int = 1
		    for i := 2; i <= n; i++ {
		        f *= i
		    }
		    return f
		}
		Execute the application from Exercise #1 again.
		Notice what's happening when you execute the application and if there is any error, solve it!
	*/

	fmt.Println(arithmetic.Factorial(5))
	/*
			Exercise #3
		Create a module (also called arithmetic) with one package (arithmetic package from the exercises above).
	*/
	/*
			Exercise #4
		Create an executable application anywhere on the disk, import the package arithmetic from GitHub and call IsPrime and Factorial functions.
	*/
	/*
			Exercise #5
		Add whatever function you like to the arithmetic package and release the module as a new major update (v2.0.0).
		Publish the major release on GitHub and then update your executable application to use the major release. Call the new function of the package in the executable application.
	*/
}
