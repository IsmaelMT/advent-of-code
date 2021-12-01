package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main() {
 
  // use os ReadLines to read line from file
  inputFile, err := os.Open("./in.txt")

  if err != nil {
    fmt.Println(err)
  }

  defer inputFile.Close()
  scanner := bufio.NewScanner(inputFile)

  var depthIncrease = 0
  var currentDepth = 0

  for scanner.Scan() {

    depth, _ := strconv.Atoi(scanner.Text())

    if currentDepth != 0 && depth > currentDepth {
      depthIncrease += 1
    } 
    currentDepth = depth
  }

  fmt.Println("Depth has increased", depthIncrease, "times")
}
