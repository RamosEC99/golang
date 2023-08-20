package basic

import (
	"fmt"
	"reflect"
)

func CastingType() {
	var i int = 42
	fmt.Println(i, reflect.TypeOf(i))
	var f float64 = float64(i)
	fmt.Println(f, reflect.TypeOf(f))
	var u uint = uint(f)
	fmt.Println(u, reflect.TypeOf(u))
}

func InferenceType() {
	var i int = 42
	fmt.Printf("i is of type %T \n", i)
	var f float64 = float64(i)
	fmt.Printf("f is of type %T \n", f)
	var u uint = uint(f)
	fmt.Printf("u is of type %T \n", u)
}
