package main

import "fmt"

type celsius float64

//type Func func()
// Instead of defining type Func as above it can be defined as a
type Func = func()

func (c celsius) String() string {
	return fmt.Sprintf("%.2f C", c)
}

func run(f Func) {
	f()
}

func hello() {
	fmt.Println("Hello")
}

// This will create same memory layout as celsius type, but type temprature will not get associated with String() func
//type temperature celsius

// To fix this we can compose celsius struct to temperature struct like
type temperature struct {
	celsius
}

// type temperature = celcius // heree temperature is of alias type of celsius and it will contain the method of celcius,
// because it is not a new type, it is same as celcius type and t = c
type anotherFunc Func

func main() {
	c := celsius(10.0)
	fmt.Println(c) // 10.00 C
	//with type temperature celsius
	// t := temperature(c)
	//fmt.Println(t) // output 10 not 10.00 C
	t := temperature{c}
	fmt.Println(t) // Now this will work and print 10.00 C
	fmt.Printf("%T\n", hello)
	run(hello)

	var f anotherFunc = hello

	fmt.Printf("Another Func Type: %T\n", f)

	//run(Func(f))

	run(f) // now this will work and you don't have to do conversion for that, type Func = func() defined like this.

}
