package main

import "fmt"

func main() {

	speed := 100 //int
	force := 2.5 //float

	//converting float to int and multiply
	res := speed * int(force)
	fmt.Println(res) //200

	speed = int(float64(speed) * force)
	fmt.Println(speed) //250

}
