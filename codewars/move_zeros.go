package main

func MoveZeros(arr []int) []int {
  var ints, zeros []int

  for _,v := range arr {
  	if v == 0 {
  		zeros = append(zeros, v)
  	} else {
  		ints = append(ints, v)
  	}
  }
  return append(ints, zeros...)
}

func main () {
	MoveZeros([]int{1,2,0,1,0,1,0,3,0,1})
}