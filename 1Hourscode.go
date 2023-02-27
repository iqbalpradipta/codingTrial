package main

import (
	"fmt"
)

/*
Score:
- Benar +1
- Salah +0
*/

func sqrt(x int) int {

	var tampung int
	/*
	   Sqrt(x)

	   Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.
	   You must not use any built-in exponent function or operator.
	   For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.

	   Example 1:

	   Input: x = 4
	   Output: 2
	   Explanation: The square root of 4 is 2, so we return 2.
	   Example 2:
	*/

	tampung += x

	tampungs := tampung / 2

	return tampungs

}

func valid_palindrome(s string) bool {
	/*
	   Valid Palindrome

	   A phrase is a palindrome if, after converting all uppercase letters into
	   lowercase letters and removing all non-alphanumeric characters,
	   it reads the same forward and backward. Alphanumeric characters include letters and numbers.
	   Given a string s, return true if it is a palindrome, or false otherwise.

	   Example 1:
	   Input: s = "A man, a plan, a canal: Panama"
	   Output: true
	   Explanation: "amanaplanacanalpanama" is a palindrome.
	*/

	for i := 0; i < len(s)/2 ; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
	
}

// func valid_anagram(s string) bool {
// 	/*
// 	   Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// 	   An Anagram is a word or phrase formed by rearranging the letters of a different
// 	   word or phrase, typically using all the original letters exactly once.

// 	   Example 1:
// 	   Input: s = "anagram", t = "nagaram"
// 	   Output: true

// 	   Example 2:
// 	   Input: s = "rat", t = "car"
// 	   Output: false
// 	*/
// }

func keyboard_row(words []string) []string {
	/*
	   Given an array of strings words, return the words that can be typed using
	   letters of the alphabet on only one row of American keyboard like the image below.

	   In the American keyboard:
	   - the first row consists of the characters "qwertyuiop",
	   - the second row consists of the characters "asdfghjkl", and
	   - the third row consists of the characters "zxcvbnm".

	   Example:
	   https://upload.wikimedia.org/wikipedia/commons/thumb/d/da/KB_United_Kingdom.svg/800px-KB_United_Kingdom.svg.png

	   Example 1:
	   Input: words = ["Hello","Alaska","Dad","Peace"]
	   Output: ["Alaska","Dad"]

	   Example 2:
	   Input: words = ["omk"]
	   Output: []
	*/

	word1 := "qwertyuiop"
	word2 := "asdfghjkl"
	word3 := "zxcvbnm"

	// kata1 := append(words, word1)
	// kata2 := append(words, word2)
	// kata3 := append(words, word3)

	var hasil []string

	hasil = append(hasil, word1, word2, word3)
	for i := 0; i < len(words)/2; i++ {
		if words[i] == word1 {
			return hasil
		} else if words[i] == word2 {
			return hasil
		} else if words[i] == word3 {
			return hasil
		}
	}

	return hasil
}

func main() {
	// task 1
	fmt.Println(sqrt(4)) // 2

	// task 2
	fmt.Println(valid_palindrome("A man, a plan, a canal: Panama")) // true

	// task 3
	// fmt.Println(valid_anagram("anagram", "nagaram")) // true

	// task 4
	fmt.Println(keyboard_row([]string{"Hello", "Alaska", "Dad", "Peace"})) // ["Alaska","Dad"]
}
