// Complete the solution so that it returns true if the first argument(string) passed in ends with the 2nd argument (also a string).

package main

import "fmt"

func solution(str string, ending string) {

	// check for empty strings
	if len(str) == 0 {
		if len(ending) == 0 {
			fmt.Printf("[+] \"%v\" ends with \"%v\"\n", str, ending)
			return
		}
		fmt.Printf("[X] \"%v\" does not end with \"%v\"\n", str, ending)
		return
	}

	if len(ending) > 0 {
		// fmt.Println("Length of str:", len(str))
		// fmt.Println("Length of ending:", len(ending))
		for i := 0; i < len(ending); i++ {
			// fmt.Println("i:", i)
			// fmt.Printf("str[%v]: %v\n", i, string(str[len(str) - len(ending) + i]))
			// fmt.Printf("ending[%v]: %v\n", i, string(ending[i]))
			currCharStr := string(str[len(str) - len(ending) + i])
			currCharEnding := string(ending[i])
			if currCharStr == currCharEnding {
				continue
			} else {
				fmt.Printf("[X] \"%v\" does not end with \"%v\"\n", str, ending)
				return
			}
		}
	}
	fmt.Printf("[+] \"%v\" ends with \"%v\"\n", str, ending)
	return
}

func main () {
	solution("", "")
	solution(" ", "")
	solution("abc", "c")
	solution("bank", "ank")
	solution("a", "z")
	solution("", "t")
	solution("$a = $b + 1", "+1")
	solution("    ", "   ")
	solution("1oo", "100")
}