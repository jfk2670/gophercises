// Given an array of integers, find the one that appears an odd number of times.
// There will always be only one integer that appears an odd number of times.

package main

func FindOdd(seq []int) int {
    count := make(map[int]int)

	for _,v := range seq {
		if count[v] == 0{
			count[v] = 1
		} else {
			delete(count, v)
		}
	}

	for n,_ := range count {
		return n
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
    //sol := [...]int{5,-1,5,10,10,1}
    for _,v := range arr {
    	FindOdd(v)
    }
}