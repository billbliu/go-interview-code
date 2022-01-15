/*
 * @Author: bill
 * @Date: 2022-01-15 16:40:05
 * @LastEditors: bill
 * @LastEditTime: 2022-01-15 16:42:01
 * @Description: go test -v 0005_str_replace_blank.go 0005_str_replace_blank_test.go
 * @FilePath: /go-interview-code/questions/0005_str_replace_blank_test.go
 */
package questions

import (
	"log"
	"testing"
)

func TestReplaceBlank(t *testing.T) {
	str := "abcDFG ssEE ffDD"

	newStr, success := replaceBlank(str)
	if success {
		log.Printf("Replace success: newStr=%s\n", newStr)
	} else {
		log.Printf("Replace failed\n")
	}

	str1 := "abcDFG22 ssEE ffDD"

	newStr, success = replaceBlank(str1)
	if success {
		log.Printf("Replace success: newStr=%s\n", newStr)
	} else {
		log.Printf("Replace failed\n")
	}
}
