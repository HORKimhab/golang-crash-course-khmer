// You can edit this code!
// Click here and start typing.
package main

// fmt to running go project
import "fmt"

/*
main() is from package main
learning Golang with HKimhab
*/
func main2() {
	// Declare multi value
	var kh, cam, gender, country string = "hkimhab", "Cambodia", "male", "Cambodia"

	// Declare without specific type 
	var a, c = 22, "Universe"
	d, e := "Factory", "1000Bitcoin"

	// Declare in block and group

	var (
		a2 int 
		b2 int = 9999
		c2 string = "Hkimhab golang"
	)

	fmt.Println("Hello Wolrd.")
	fmt.Println("Hello Universe.")
	fmt.Println("Hello Universe by HKimhab in Golang.")
	// Single Line comment
	fmt.Println("Hello Crush, nice to see you.")
	fmt.Println("Mulit declare value---")
	fmt.Println(kh)
	fmt.Println(cam)
	fmt.Println(gender)
	fmt.Println(country)
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(a2)
	fmt.Println(b2)
	fmt.Println(c2)
}
