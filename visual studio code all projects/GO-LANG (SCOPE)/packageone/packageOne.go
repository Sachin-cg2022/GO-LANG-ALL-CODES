package packageone

// package variable exporting
var privateVar = " i am private"          //"private applicable into inside package"
var PublicVar = "i am public or exported" //"accessable outside the package"//
// P should br capital for public

func notExported() { // avilable only inside packageone

}

func Exported() { // accessable outside the package

}
