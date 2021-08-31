//Sample program to look at pointers in go. They are a lot like c++ (YAY)

package main

import(
	"fmt" 
)

func main() {
	var i int = 10; //Initialize the integer

	var ptrToI *int = &i; //Make a pointer to the address of i using &

	fmt.Println(*ptrToI); //Dereference the pointer using * and print
}