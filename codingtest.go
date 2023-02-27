package main

import (
	"fmt"
	"strings"
)

/*
Score:
- Benar +1
- Salah +0
*/

func sqrt(x int) int {
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

	if x == 0 {
        return 0
    }

    l, r := 1, x
    for l <= r {
        m := (l + r) / 2
        if m == x/m {
            return m
        } else if m < x/m {
            l = m + 1
        } else {
            r = m - 1
        }
    }

    return r
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

func valid_anagram(s string) bool {
	/*
	   Given two strings s and t, return true if t is an anagram of s, and false otherwise.
	   An Anagram is a word or phrase formed by rearranging the letters of a different
	   word or phrase, typically using all the original letters exactly once.

	   Example 1:
	   Input: s = "anagram", t = "nagaram"
	   Output: true

	   Example 2:
	   Input: s = "rat", t = "car"
	   Output: false
	   
	*/
	var t string

	if len(s) != len(t) {
        return false
    }

    count := make(map[rune]int)
    for _, v := range s {
        count[v]++
    }

    for _, v := range t {
        count[v]--
        if count[v] < 0 {
            return false
        }
    }

    return true
}

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

	rows := map[rune]int{
        'q': 1, 'w': 1, 'e': 1, 'r': 1, 't': 1, 'y': 1, 'u': 1, 'i': 1, 'o': 1, 'p': 1,
        'a': 2, 's': 2, 'd': 2, 'f': 2, 'g': 2, 'h': 2, 'j': 2, 'k': 2, 'l': 2,
        'z': 3, 'x': 3, 'c': 3, 'v': 3, 'b': 3, 'n': 3, 'm': 3,
    }
    
    var hasil []string
    for _, word := range words {
        row := 0
        SingleRow := true
        for _, v := range strings.ToLower(word) {
            if row == 0 {
                row = rows[v]
            } else if rows[v] != row {
                SingleRow = false
                break
            }
        }
        if SingleRow {
            hasil = append(hasil, word)
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
	fmt.Println(valid_anagram("anagram", "nagaram")) // true

	// task 4
	fmt.Println(keyboard_row([]string{"Hello", "Alaska", "Dad", "Peace"})) // ["Alaska","Dad"]
}
