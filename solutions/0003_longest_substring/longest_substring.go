package longestsubstring

import (
	"fmt"
)

/*
Given a string s, find the length of the longest substring without duplicate characters.



Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.


Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	currentIndex := 0
	currentStringMap := map[string]bool{}
	fmt.Println("maxLength", maxLength)
	for {
		if currentIndex >= len(s) {
			break
		}
		mapKey := fmt.Sprintf("%c_%d", s[currentIndex], currentIndex)
		fmt.Println("mapKey", mapKey)

		if _, exists := currentStringMap[mapKey]; exists {
			maxLength = max(maxLength, len(currentStringMap))
			for key := range currentStringMap {
				fmt.Println("key", key)
				delete(currentStringMap, key)
				if key == mapKey {
					break
				}
			}
			currentStringMap[mapKey] = true
			currentIndex++
			fmt.Println("currentStringMap", currentStringMap)
			continue
		}

		currentStringMap[mapKey] = true
		currentIndex++
		maxLength = max(maxLength, len(currentStringMap))
	}

	return maxLength
}
