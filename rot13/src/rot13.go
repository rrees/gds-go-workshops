package main

import "fmt"
import "bufio"
import "os"
import "strings"

func main() {

  reader := bufio.NewReader(os.Stdin)

  text, _ := reader.ReadString('\n')

  encoded := strings.Map(Rot13, text)


  fmt.Println(encoded)

}

func Rot13(r rune) rune {
  if(IsAlpha(r)) {
    return r + 1
  } else {
    return r
  }
}

func IsAlpha(r rune) bool {
  return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}
