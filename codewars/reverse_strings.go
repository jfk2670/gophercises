// Write a function that takes in a string of one or more words, and returns the same string, 
// but with all five or more letter words reversed. 
// Strings passed in will consist of only letters and spaces.
// Spaces will be included only when more than one word is present.

package main

import "strings"

func SpinWords(str string) string {
  split := strings.Split(str, " ")

  for index,word := range split {
  	reversed := ""
  	if len(word) >= 5 {
  		for i := len(word) - 1; i >= 0; i-- {
  			reversed += string(word[i])
  	 	}
  	 split[index] = reversed
  	}
  }

  return strings.Join(split, " ")
}

func main () {
	SpinWords("Welcome")
    SpinWords("to")
    SpinWords("CodeWars")
    SpinWords("Hey fellow warriors")
    SpinWords("Burgers are my favorite fruit")
    SpinWords("Pizza is the best vegetable")
}