//This is a package file, equivalent of .h
//note that any func/variable starting with UPPER CASE is public, lower case is private!!

package subtractor /*name the package, this can be a sub-package or the same
					 as another package name from another file!!*/

func Subtract (i, j int)(int){ //Upper case for external!
	
	var total int;
	
	total = i - j;

	return total;
}