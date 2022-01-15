/*
 * @Author: bill
 * @Date: 2022-01-15 16:37:26
 * @LastEditors: bill
 * @LastEditTime: 2022-01-15 16:42:50
 * @Description: go test -v 0005_str_replace_blank.go 0005_str_replace_blank_test.go
 * @FilePath: /go-interview-code/questions/0005_str_replace_blank.go
 */
package questions

import (
	"strings"
	"unicode"
)

func replaceBlank(s string) (string, bool) {
	if len([]rune(s)) > 1000 {
		return s, false
	}

	for _, v := range s {
		if string(v) != " " && !unicode.IsLetter(v) {
			return s, false
		}
	}

	return strings.Replace(s, " ", "%20", -1), true
}
