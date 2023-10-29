// Your goal in this kata is to implement a difference function, which subtracts one list from another and returns the result.
// It should remove all values from list a, which are present in list b keeping their order.
// If a value is present in b, all of its occurrences must be removed from the other:

package main

import "fmt"

func ArrayDiff(a, b []int) []int {
  	var newArray []int
  	var found bool = false

  	// loop through a
  	for a_i := 0; a_i < len(a); a_i++ {
  		fmt.Printf("a[%v]: %v\n", a_i, a[a_i])

  		found = false
  		//loop through b
  		for b_i := 0; b_i < len(b); b_i++ {
  			fmt.Printf("b[%v]: %v\n", b_i, b[b_i])
  			if b[b_i] == a[a_i] {
  				found = true
  			}
  		}

  		if found == false {
  			newArray = append(newArray, a[a_i])
  		}
  	}
  	fmt.Println("newArray:", newArray)
  	return newArray
}

func main () {
	var emptyArr []int
	ArrayDiff([]int{1,2},[]int{1})
	ArrayDiff([]int{1,2,2},[]int{1})
	ArrayDiff([]int{1,2,2},[]int{2})
	ArrayDiff([]int{1,2,2},emptyArr)
	ArrayDiff(emptyArr,[]int{1,2})
	ArrayDiff([]int{1,2,3},[]int{1,2})
}