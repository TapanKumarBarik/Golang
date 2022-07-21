package main

import (
	"fmt"
	"path"
)

var start int

func main() {

	//type 1
	var dir, filename string
	dir, filename = path.Split("css/test/main.css")

	fmt.Println("dir : ", dir)
	fmt.Println("filename : ", filename)

	//type 2
	var filename2 string
	_, filename2 = path.Split("css/test.css")
	fmt.Println("filename2 : ", filename2)

	//type 3
	//no need to initilize the varible

	_, filename3 := path.Split("css/test3.css")
	fmt.Println("filename3 : ", filename3)

	//type 4
	//
	filepath, filename4 := path.Split("css/test4.css")
	fmt.Println("filename4 : ", filename4)
	fmt.Println("filepath : ", filepath)

	//type 4
	//declare variable in a group
	var (
		video    string
		duration int
		current  int
	)

	fmt.Println(video)
	fmt.Println(duration)
	fmt.Println(current)

	//if we know the value user short hand
	width, height := 40, 50
	fmt.Println(width)
	fmt.Println(height)
}
