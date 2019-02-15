package main

import (
  "fmt"
  "time"
  "math/rand"
)

func main() {
  newSeed()

  len := 0
  fmt.Printf("length of slice to be sorted: ")
  fmt.Scanf("%d", &len)

  numbers := randslice(len)
  fmt.Println("original:", numbers)

  start := time.Now()
  sorted := 0

  for ; sorted < len - 1; {
    sorted = 0

    for i := 1; i < len; i ++ {
      if numbers[i] < numbers[i - 1] {
        swap(numbers, i, i - 1)
      } else {
        sorted++
      }
    }
  }

  fmt.Println("sorted:  ", numbers)
  elapsed := time.Since(start)
  fmt.Println(elapsed)
}

func randslice(len int) []int {
  var numbers []int

  for i := 0; i < len; i++ {
    numbers = append(numbers, rand.Intn(10))
  }

  return numbers
}

func swap(numbers []int, ia int, ib int)  {
  temp := numbers[ia]
  numbers[ia] = numbers[ib]
  numbers[ib] = temp
}

func newSeed()  {
  rand.Seed(time.Now().UnixNano())
}
