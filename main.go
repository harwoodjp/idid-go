package main

import (
  "os"
  "idid/models"
)

func main() {
  args := os.Args[1:]

  if len(args) == 1 {
    journal := models.Journal{FilePath:"/Users/justin.harwood/Documents/.ididjournal" }

    if args[0] == "what" {
      journal.Read()
      os.Exit(0)
    }

    journal.Write(args[0])
  }

}
