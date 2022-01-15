/*
 * @Author: bill
 * @Date: 2022-01-15 16:15:19
 * @LastEditors: bill
 * @LastEditTime: 2022-01-15 16:17:32
 * @Description: go test -v 0004_str_is_regroup.go 0004_str_is_regroup_test.go
 * @FilePath: /go-interview-code/questions/0004_str_is_regroup.go
 */
package questions

import "strings"

func isRegroup(s1, s2 string) bool {
	s1Len := len([]rune(s1))
	s2Len := len([]rune(s2))

	if s1Len > 5000 || s2Len > 5000 || s1Len != s2Len {
		return false
	}

	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}

	return true
}
