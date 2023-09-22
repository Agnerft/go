package main

import (
	"fmt"
	"testeTela/data"
	"testeTela/screen"
)

func main() {
	screen.ScreenPrincipal()

	fmt.Printf(data.GetInfo1())
}
