package ransom_note

func CanConstruct(ransomNote string, magazine string) bool {
	ransomNoteLength := len(ransomNote)
	ransomNoteCount := 0
	usedMagazine := make(map[int]bool)

	for i := 0; i < ransomNoteLength; i++ {
		for j := 0; j < len(magazine); j++ {
			if ransomNote[i] == magazine[j] && !usedMagazine[j] {
				ransomNoteCount++
				usedMagazine[j] = true
				break
			}
		}
	}

	return ransomNoteCount == ransomNoteLength
}
