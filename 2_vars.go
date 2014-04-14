// Ruby:

// arr = []

// 5.times do |n|
//   arr << n
// end

// puts b

package main

import "fmt"

func main() {

  // bool
  // string
  // int  int8  int16  int32  int64
  // uint uint8 uint16 uint32 uint64 uintptr
  // byte // alias for uint8
  // rune // alias for int32
  //      // represents a Unicode code point
  // float32 float64
  // complex64 complex128

	// implicit type
	s := "Hello World!"

	// explicit type
	var arr []int

	for n := 0; n < 5; n++ {
		arr = append(arr, n)
	}

	fmt.Println(s)
	fmt.Println(arr)
}
