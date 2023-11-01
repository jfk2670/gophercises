package main

import "fmt"

func MoveZeros(arr []int) []int {
  newArr := []int

  for i,v := range arr {
  	if v == 0 {
  		newArr = append(newArr, v)
  	}
  }
  fmt.Println(newArr)
  return 0
}

func main () {
	MoveZeros([]int{1,2,0,1,0,1,0,3,0,1})
}