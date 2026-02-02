package main

import (
	"fmt"
	"strings"
)

func uniqueS(str string) bool {
	lstr := strings.ToLower(str)

	hash := make(map[rune]struct{})

	for _, s := range lstr{
		if _, ok := hash[s]; ok {
			return false
		}
		hash[s] = struct{}{}
	}
	return true
}

func main(){
	fmt.Println(uniqueS("abcd"))
	fmt.Println(uniqueS("abCdefAaf"))
	fmt.Println(uniqueS("aabcd"))
}
