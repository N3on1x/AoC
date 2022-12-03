// This is part 2 of Advent of code 2022
package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

/* Making a min heap */
type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i,j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
  *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
  old := *h
  n := len(old)
  x := old[n-1]
  *h = old[0 : n-1]
  return x
}

/* Main function */
func main() {
  // Timing
  start0 := time.Now()

  // Read and open filename from arguments
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

  // Start of algorithm
  s := bufio.NewScanner(file)

  hp := &IntHeap{0,0,0}
  heap.Init(hp)

  heapFixCount := 0

  cur := 0
  for s.Scan() {
    line := s.Text()
    if line == "" {
      if cur > (*hp)[0] {
        (*hp)[0] = cur
        heap.Fix(hp,0)
        heapFixCount++
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

  sum := 0
  for _, num := range (*hp) {
    sum += num
  }
  fmt.Println("The three elves carrying the most Calories is carrying a total of", sum, "Calories")
  fmt.Println("Total time: ", time.Since(start0))
  fmt.Println("Number of heap updates:", heapFixCount)
}
