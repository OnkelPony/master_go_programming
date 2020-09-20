package main

/*
multiline comment
is used for debugging
mainly
*/

// line comment is used more often

var s string
var i int
var msg string
var val float64
var err error
var done bool

var maxValue = 666  // OK - camelCase
var max_value = 666 // NOK - snake_case

func main() {
	writeToDB = true  // OK - acronym in all caps
	writeToDb = false // NOK
}
