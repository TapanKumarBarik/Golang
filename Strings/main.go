package main

import "fmt"

func main() {

	fmt.Println(`hello`)

	s := "how are you"

	fmt.Println(s)

	s = `<html>
			<head>
				<title>new</title>
			</head>
			<body>
				<h1>This is something</h1>
			</body>
		</html>`

	fmt.Println(s)

	//escape characters

	s = "c:\\project\\css\\main.css"
	fmt.Println(s)

	s = `c:\project\css\main.css`

	fmt.Println(s)

	//tab
	//\t
	//new line
	//\n

}
