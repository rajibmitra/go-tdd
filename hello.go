package main

import "fmt"

func hello() string {
	return "hello Rajib"
}

func main() {
	fmt.Println(hello())
	fmt.Println(helloHindi())
}

func helloHindi() string {
	return "हल्लो रजिब"
}
