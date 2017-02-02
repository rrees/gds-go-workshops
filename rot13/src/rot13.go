package main

import "fmt"
import "bufio"
import "os"

func main() {
  fmt.Println("Hello from the encoder")

  reader := bufio.NewReader(os.Stdin)

  text, _ := reader.ReadString('\n')

  fmt.Println(text)
}
