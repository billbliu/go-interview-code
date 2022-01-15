/*
 * @Author: bill
 * @Date: 2022-01-15 15:00:43
 * @LastEditors: bill
 * @LastEditTime: 2022-01-15 15:26:27
 * @Description: go test -v 0002_is_no_repeat_string.go 0002_is_no_repeat_string_test.go
 * @FilePath: /go-interview-code/questions/0002_is_no_repeat_string_test.go
 */
package questions

import (
	"log"
	"testing"
)

func TestIsNoRepeatString(t *testing.T) {
	str := "akjce1834opw"
	log.Println(isNoRepeatString1(str))
	log.Println(isNoRepeatString2(str))

	str1 := "akjce1834kopwaaaa"
	log.Println(isNoRepeatString1(str1))
	log.Println(isNoRepeatString2(str1))
}
