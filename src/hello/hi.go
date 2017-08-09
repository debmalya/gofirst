package main

import "fmt"

func main() {
    fmt.Printf("What's your name?");
    var name string
    fmt.Scanf("%s",&name)
    fmt.Printf("hello "+ name)
}
