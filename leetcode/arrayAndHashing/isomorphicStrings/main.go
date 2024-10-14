package main 

import(
	"fmt"
)

func isIsomorphic(s, t string) bool {
	if len(s) != len(t) {
		return false 
	}

	sToT := make(map[byte]byte)
	tToS := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		charS := s[i]
		charT := t[i]

		if val, ok := sToT[charS]; ok {
			if val != charT {
				return false 
			}
		} else {
			sToT[charS] = charT 
		}

		if val, ok := tToS[charT]; ok {
			if val != charS {
				return false
			}
		} else {
			tToS[charT] = charS 
		}
	}
	return true 
}

func main() {
	s := "egg"
    t := "add"
    fmt.Println(isIsomorphic(s, t)) 
}