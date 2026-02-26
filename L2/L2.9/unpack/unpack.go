package unpack

import (
	"errors"
	"strings"
)

var errFirstDigit error = errors.New("string cannot start with a digit")

func Unpack(str string) (string, error) {
	if str == "" {
		return "", nil
	}
	if str[0] <= '9' {
		return "", errFirstDigit
	}

	var (
		prev rune
		num  int
		hasNum bool
		res  strings.Builder
	)

	for _, char := range str {
		if char <= '9' {
			if prev == 92 {
				prev = char
				continue
			}
			hasNum = true
			num *= 10
			num += int(char - '0')
		} else {
			if hasNum {
				res.WriteString(strings.Repeat(string(prev), num))
				hasNum = false
				num = 0
			} else if prev != 0 {
				res.WriteRune(prev)
			}
			prev = char
		}
	}

	if hasNum {
		res.WriteString(strings.Repeat(string(prev), num))
	} else if prev != 0 {
		res.WriteRune(prev)
	}
	
	return res.String(), nil
}