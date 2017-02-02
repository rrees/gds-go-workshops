package main

import (
  "bufio"
  "os"
  "io"
  "fmt"
  "log"
)

func main() {
  reader := bufio.NewReader(os.Stdin)

  fmt.Println("Reading stream")

  result:= [3]int{0,0,0}

  idx := 0

  for {
    b, err := reader.ReadByte()

    if err == io.EOF {
      break
    } else if err != nil {
      log.Fatalf("Error reading stream: %s", err)
    }

    if b >= 30 && b <= 39 {
      fmt.Println(b)
      result[idx] = int(b - 30)
      fmt.Println(sum(result) % 10)

      idx = (idx + 1) % 3

    }
  }

  fmt.Println("Stream read!")
}

func sum(a [3]int) int {
  sum := 0

  for _, v := range a {
    sum += v
  }

  return sum
}
