package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "unicode"

func main() {

  reader := bufio.NewReader(os.Stdin)

  text, _ := reader.ReadString('\n')

  encoded := strings.Map(rot13, text)


  fmt.Println(encoded)

}

func rot13(r rune) rune {
  return unicode.ToUpper(r)
}
