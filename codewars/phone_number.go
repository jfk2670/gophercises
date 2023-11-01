package main

import "regexp"
import "fmt"
import "strconv"

func CreatePhoneNumber(numbers [10]uint) string {
	number := "("

	for i,v := range numbers {
		m,_ := regexp.MatchString("[0-9]", strconv.FormatUint(uint64(v), 10))
		if m {
			if i == 3 {
				number += ") "
			}
			if i == 6 {
				number += "-"
			}
			number += strconv.FormatUint(uint64(v), 10)
		} else {
			return "Non-int provided"
		}
	}
	fmt.Println(number)
	return number
}

func main () {
	CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})
}