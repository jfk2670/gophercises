// Complete the solution so that it returns true if the first argument(string) passed in ends with the 2nd argument (also a string).

package main

func solution(str string, ending string) {

	// check for empty strings
  if len(str) == 0 || len(ending) > len(str) {
		if len(ending) == 0 {
			return true
		}
		return false
	}

	if len(ending) > 0 {
		for i := 0; i < len(ending); i++ {
			currCharStr := string(str[len(str) - len(ending) + i])
			currCharEnding := string(ending[i])
			if currCharStr == currCharEnding {
				continue
			} else {
				return false
			}
		}
	}
	return true
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