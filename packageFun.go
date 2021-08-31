//Lets do #include *cough* I mean packages!
//Go gets mad if we import packages and don't use them.

package main //main is a special package that indicates an executable

import ( //Here is #include's equivalent.
"fmt" //pre-added libs are typically in goroot already. no abs pathing required.
"practice/packages/math" //Lets add one of our own libraries too. This one does need a path!
)

func main() {
	var i, j int;

	i = 8;
	j = 2;

	fmt.Println(subtractor.Subtract(i, j)); //Use a function from the subtract package.
}

