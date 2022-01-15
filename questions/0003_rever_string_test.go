/*
 * @Author: bill
 * @Date: 2022-01-15 15:35:32
 * @LastEditors: bill
 * @LastEditTime: 2022-01-15 15:38:44
 * @Description: go test -v 0003_rever_string.go 0003_rever_string_test.go
 * @FilePath: /go-interview-code/questions/0003_rever_string_test.go
 */
package questions

import (
	"log"
	"testing"
)

func TestReverString(t *testing.T) {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	revStr, isSuccess := reverString(str)
	if isSuccess {
		log.Printf("success: revstr=%s\n", revStr)
	} else {
		log.Printf("failed: revstr=%s\n", revStr)
	}

	str1 := "aaaaaaabbbbbbbbbbbbccccccccddddddd"

	revStr1, isSuccess := reverString(str1)
	if isSuccess {
		log.Printf("success: revstr=%s\n", revStr1)
	} else {
		log.Printf("failed: revstr=%s\n", revStr1)
	}
}
