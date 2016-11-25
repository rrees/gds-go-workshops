package main

import "fmt"
import "net/http"
import "log"
import "math/rand"
import "encoding/json"

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

  diceResult := rollDice(die)

  log.Print("Die value is %d", diceResult)

  result := RollResult{dieType, diceResult}

  log.Print(result)

  jsonPayload, error := json.Marshal(result)

  if error != nil {
    log.Print("Failed to marshal the payload")
    http.Error(w, fmt.Sprintf("Failed to create the JSON payload %v", error), http.StatusInternalServerError)
    return
  }

  log.Print(jsonPayload)

  w.Header().Set("Content-Type", "application/json")
  w.Write(jsonPayload)
}

type RollResult struct {
  Dice string
  Result int
}
