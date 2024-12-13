package annagram

import (
	"strings"
)

func anagram(s1, s2 string) bool {
	s1map := make(map[rune]int)
	s2map := make(map[rune]int)
	getRuneMap := func(s string, sMap map[rune]int) {
		for _, r := range []rune(strings.ReplaceAll(strings.ToLower(s), " ", "")) {
			sMap[r]++
		}
	}
	getRuneMap(s1, s1map)
	getRuneMap(s2, s2map)
	// fmt.Println(s1map)
	// fmt.Println(s2map)

	len1 := len(s1map)
	len2 := len(s2map)
	if len1 != 0 && len2 != 0 {
		if len1 != len2 {
			// fmt.Printf("s1map size =%d, s2map size =%d", len1, len2)
			return false
		} else {
			for key := range s1map {
				if _, isExist := s2map[key]; !isExist {
					return false
				}
			}
			return true
		}
	} else {
		// fmt.Printf("s1map size =%d, s2map size =%d", len1, len2)
		return false
	}
}

// func main() {
// 	fmt.Println(anagram("testStr1", "testStr1"))
// 	fmt.Println(anagram("testStr1", "testStr2"))
// }
