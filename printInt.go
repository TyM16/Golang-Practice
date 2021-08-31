//Sample program for printing a sum in go.

package main //Package name for importing elsewhere

import "fmt" // essentially #include in go

func main() { //func with no return has no return type, do not add void. return type comes after declaration of func.
	var i, j int; // variable. Goes type (var, func, etc.), name, type of variable.

	i = 4;
	j = 6;

	var output int = trivialVariableAdder(i, j);

	fmt.Println(output); // equivalent of cout in go.
};

func trivialVariableAdder(i, j int) (int) { //Note functions can have multiple returns in go!!
	var sum int;

	sum = i+j;

	return sum;
}