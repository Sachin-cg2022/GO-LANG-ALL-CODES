package packageone

import "fmt"

var PackageVar = ", inside packageone package level variable"

func PrintMe(myVar, blockVar string) {

	fmt.Println(myVar, blockVar, PackageVar)

}
