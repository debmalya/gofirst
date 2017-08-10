package main


import (
  "bufio"
  "os"
  "fmt"
)

func main() {
    fmt.Printf("What's your name?");
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() // use `for scanner.Scan()` to keep reading
    name := scanner.Text()
    fmt.Printf("hello "+ name)
    fmt.Printf("\n")
}
