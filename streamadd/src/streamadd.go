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

  for {
    b, err := reader.ReadByte()

    if err == io.EOF {
      break
    } else if err != nil {
      log.Fatalf("Error reading stream: %s", err)
    }

    if b >= 30 && b <= 39 {
      fmt.Println(b)
    }
  }

  fmt.Println("Stream read!")
}
