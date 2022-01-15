/*
 * @Author: bill
 * @Date: 2022-01-15 11:28:41
 * @LastEditors: bill
 * @LastEditTime: 2022-01-15 12:08:50
 * @Description: go test -v 0001_print_letter_number.go 0001_print_letter_number_test.go
 * @FilePath: /go-interview-code/questions/0001_print_letter_number.go
 */
package questions

import (
	"fmt"
	"strings"
	"sync"
)

func printLetterNumber() {
	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}

	// print number
	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
				break
			default:
				break
			}
		}
	}()
	wait.Add(1)

	// print letter
	go func(wait *sync.WaitGroup) {
		letterStr := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-letter:
				if i >= strings.Count(letterStr, "")-1 {
					wait.Done()
					return
				}
				fmt.Print(letterStr[i : i+1])
				i++
				// 防止访问越界数组
				if i >= strings.Count(letterStr, "") {
					i = 0
				}
				fmt.Print(letterStr[i : i+1])
				i++
				number <- true
				break
			default:
				break
			}
		}
	}(&wait)

	// 开始打印
	number <- true
	wait.Wait()
}
