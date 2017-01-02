package main

import (
	"fmt"
)

func getNameByUser() string {
	var name string
	fmt.Printf("What's your name? ")
	if n, _ := fmt.Scanln(&name); n <= 0 {
		return "you just have pressed with empty string. Please press enter key after typed something"
	} else {
		return name
	}
}

func makeText(text string) string{
        return fmt.Sprintf("%s has %d characters" , text , len(text))
}

func printText( text string ){
        fmt.Println(text)
}

func main() {
	printText(makeText(getNameByUser()))
}
