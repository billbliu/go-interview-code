/*
 * @Author: bill
 * @Date: 2022-01-15 15:32:04
 * @LastEditors: bill
 * @LastEditTime: 2022-01-15 15:38:48
 * @Description: go test -v 0003_rever_string.go 0003_rever_string_test.go
 * @FilePath: /go-interview-code/questions/0003_rever_string.go
 */
package questions

func reverString(s string) (string, bool) {
	str := []rune(s)
	lenght := len(str)

	if lenght > 5000 {
		return s, false
	}

	for i := 0; i < lenght/2; i++ {
		str[i], str[lenght-1-i] = str[lenght-1-i], str[i]
	}
	return string(str), true
}
