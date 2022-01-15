package main

import "fmt"

type AccountNumbers struct {
	AccountNumber [9]Display
}

type Display struct {
	DisplayRow1 [3]string
	DisplayRow2 [3]string
	DisplayRow3 [3]string
}

func main() {
	fmt.Println("hola mundo")
}
