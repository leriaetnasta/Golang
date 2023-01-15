package main

import "fmt"

func main() {
	i := 1
	fmt.Println(i)
	//conversion
	var j float32 = 3.5
	fmt.Printf("%v", j)
	i = int(j)
	fmt.Println(i)

	fmt.Println(i)
}
