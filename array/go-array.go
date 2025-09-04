// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// fmt to running go project

/*
main() is from package main
learning Golang with HKimhab
Syntax of Array Go
- var array_name = [length]datatype{values} // here length is defined

- var array_name = [...]datatype{values} // here length is inferred

--- With := sign
- array_name := [length]datatype{values} // here length is defined

- array_name := [...]datatype{values} // here length is inferred
*/
func main() {
	fmt.Println("Learning Go array now")
	var array1 = [3]string{"Bitcoin", "Universe", "Linux"}

	var array1_1 = [...]string{"Bitcoin", "HKimhab", "Linux", "Kali Linux"}

	// array_name := [length]datatype{values} // here length is defined
	array2 := [2]int{1993, 2000}

	array2_2 := [...]int{1993, 2000, 487294}

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println("--- Array interface ---")
	fmt.Println(array1_1)
	fmt.Println(array2_2)
}
