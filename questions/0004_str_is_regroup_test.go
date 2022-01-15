/*
 * @Author: bill
 * @Date: 2022-01-15 16:18:37
 * @LastEditors: bill
 * @LastEditTime: 2022-01-15 16:26:30
 * @Description: go test -v 0004_str_is_regroup.go 0004_str_is_regroup_test.go
 * @FilePath: /go-interview-code/questions/0004_str_is_regroup_test.go
 */
package questions

import (
	"log"
	"testing"
)

func TestStrIsRegGroup(t *testing.T) {
	s1 := "aabbccdd"
	s2 := "bbccddaa"
	s3 := "bbccdaa"

	isReg := isRegroup(s1, s2)
	log.Printf("%v\n", isReg)

	isReg1 := isRegroup(s1, s3)
	log.Printf("%v\n", isReg1)
}
