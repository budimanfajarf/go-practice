package ransom_note

func CanConstruct(ransomNote string, magazine string) bool {
	// Approach 1: Using map to track used characters in magazine
	// ransomNoteLength := len(ransomNote)
	// ransomNoteCount := 0
	// usedMagazine := make(map[int]bool)

	// for i := 0; i < ransomNoteLength; i++ {
	// 	for j := 0; j < len(magazine); j++ {
	// 		if ransomNote[i] == magazine[j] && !usedMagazine[j] {
	// 			ransomNoteCount++
	// 			usedMagazine[j] = true
	// 			break
	// 		}
	// 	}
	// }

	// return ransomNoteCount == ransomNoteLength
	// time complexity: O(n*m)
	// space complexity: O(m)

	// Approach 2: Using character count map
	charCount := make(map[rune]int)

	for _, char := range magazine {
		charCount[char]++
	}

	for _, char := range ransomNote {
		if charCount[char] == 0 {
			return false
		}
		charCount[char]--
	}

	return true
	// time complexity: O(m)
	// space complexity: O(k)
}
