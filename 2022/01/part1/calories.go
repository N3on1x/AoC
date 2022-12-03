package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
  start0 := time.Now()

  if len(os.Args) < 2 {
    log.Fatal("No file name was provided.")
  }
  args := os.Args
  fname := args[1]
  file, err := os.Open(fname)
  defer file.Close()
  if err != nil {
    log.Fatal(err)
  }

  s := bufio.NewScanner(file)
  max, cur := 0, 0
  for s.Scan() {
    line := s.Text()
    if line == "" {
      if cur > max {
        max = cur
      }
      cur = 0
      continue
    }
    num, err := strconv.ParseInt(line, 10, 0)
    if err != nil {
      log.Fatal(err)
    }
    cur += int(num)
  }
  fmt.Println("The elf carrying the most Calories is carrying", max, "Calories")
  fmt.Println("Total time: ", time.Since(start0))
}
