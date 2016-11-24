package main

import "fmt"
import "net/http"
import "log"
import "math/rand"

func main() {
  fmt.Println("Hello world")

  http.HandleFunc("/roll", diceRoller)

  http.ListenAndServe(":6001", nil)
}

func determineDie(userSuppliedDie string, defaultDie string) string {
  if userSuppliedDie == "" {
    return defaultDie
  } else {
    return userSuppliedDie
  }
}

func parseDie(diceString string) (int, error) {
  return 6, nil
}

func rollDice(maxValue int) int {
  return rand.Intn(maxValue) + 1
}

func diceRoller(w http.ResponseWriter, r *http.Request) {
  dieTypeParam := r.FormValue("die")

  dieType := determineDie(dieTypeParam, "D6")

  die, error := parseDie(dieType)

  if error != nil {
    http.Error(w, fmt.Sprintf("Invalid die %s", dieTypeParam), http.StatusBadRequest)
    log.Print(error)
    return
  }

  fmt.Fprintf(w, "Die value is %d", rollDice(die))
}
