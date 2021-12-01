
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

  var slidingWindow[3] int
  var currentIndex = 0

  for scanner.Scan() {

    depth, _ := strconv.Atoi(scanner.Text())
    
    slidingWindow[currentIndex % 3] = depth

    if currentIndex >= 2 {
      
      slidingWindowDepth := slidingWindow[0] + slidingWindow[1] + slidingWindow[2]
      fmt.Println(slidingWindow[0], slidingWindow[1], slidingWindow[2])

      if currentDepth != 0 && slidingWindowDepth > currentDepth {
        depthIncrease++
      }

      currentDepth = slidingWindowDepth
    }   
    currentIndex++
  }

  fmt.Println("Depth has increased", depthIncrease, "times")
}
