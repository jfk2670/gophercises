// Given an array of integers, find the one that appears an odd number of times.
// There will always be only one integer that appears an odd number of times.

package main

import "fmt"

func FindOdd(seq []int) int {
    count := make(map[int]int)

	for _,v := range seq {
		if count[v] == 0{
			//fmt.Println("Value not in count, adding...")
			count[v] = 1
		} else {
			//fmt.Println("Value in count, increasing...")
			count[v] += 1
		}
	}

	for n,c := range count {
		if c % 2 == 1 {
			return n
		}
	}

    return 0
}

func main () {
	arr := [...][]int{
        {20,1,-1,2,-2,3,3,5,5,1,2,4,20,4,-1,-2,5},
        {1,1,2,-2,5,2,4,4,-1,-2,5},
        {20,1,1,2,2,3,3,5,5,4,20,4,5},
        {10},
        {1,1,1,1,1,1,10,1,1,1,1},
        {5,4,3,2,1,5,4,3,2,10,10},
    }
    sol := [...]int{5,-1,5,10,10,1}
    for i,v := range arr {
        fmt.Sprintf("Testing input %v",v)
        fmt.Println("Answer should be:", sol[i])
        fmt.Println(FindOdd(v))
    }
}