package models

import (
  "fmt"
  "io"
  "os"
  "log"
  "time"
  "github.com/fatih/color"
)

type Journal struct {
  FilePath string
}


func (j *Journal) Write(message string) {
  newEntry := Entry{time.Now(), message}
  f, err := os.OpenFile(j.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

  if err != nil {
    log.Fatal(err)
  }

  if _, err := f.Write([]byte(newEntry.ToString()+"\n")); err != nil {
    log.Fatal(err)
  }

  if err := f.Close(); err != nil {
    log.Fatal(err)
  }
}

func (j *Journal) Read() {
  f, err := os.Open(j.FilePath)
  if err != nil {
    fmt.Println(err)
  }

  var buff []byte
  defer f.Close()
  buff = make([]byte, 1024)
  for {
    n, err := f.Read(buff)

    if n > 0 {
      color.Green(string(buff[:n-1]))
    }

    if err == io.EOF {
      break
    }
  }
}
