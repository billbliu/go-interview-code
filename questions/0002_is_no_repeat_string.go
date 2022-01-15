/*
 * @Author: bill
 * @Date: 2022-01-15 15:00:13
 * @LastEditors: bill
 * @LastEditTime: 2022-01-15 15:22:50
 * @Description: go test -v 0002_is_no_repeat_string.go 0002_is_no_repeat_string_test.go
 * @FilePath: /go-interview-code/questions/0002_is_no_repeat_string.go
 */
package questions

import "strings"

func isNoRepeatString1(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}

	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}

	return true
}

func isNoRepeatString2(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}

	for k, v := range s {
		if v > 127 {
			return false
		}
		if strings.Index(s, string(v)) != k {
			return false
		}
	}

	return true
}
