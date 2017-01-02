package main

import (
	"fmt"
)

func getNameByUser() string{
        var str string
        fmt.Printf("What's your name? ")
	fmt.Scanf("%s", &str)
        return str
}

func makeText( text string ) string{
        return "Hello, " + text + ", nice to meet you" 
}

func printText( text string ){
        fmt.Println(text)
}
func main() {
	fmt.Println("You will see Hello World message with function called")
        printText(makeText(getNameByUser()))
}
