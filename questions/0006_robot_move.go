/*
 * @Author: bill
 * @Date: 2022-01-15 16:50:15
 * @LastEditors: bill
 * @LastEditTime: 2022-01-15 17:04:12
 * @Description: go test -v 0006_robot_move.go 0006_robot_move_test.go
 * @FilePath: /go-interview-code/questions/0006_robot_move.go
 */
package questions

import "unicode"

const (
	Left = iota
	Top
	Right
	Bottom
)

func robotMove(cmd string, x0 int, y0 int, z0 int) (x, y, z int) {
	x, y, z = x0, y0, z0
	repeat := 0
	repeatCmd := ""

	for _, s := range cmd {
		switch {
		case unicode.IsNumber(s):
			repeat = repeat*10 + (int(s) - '0')
		case s == ')':
			for i := 0; i < repeat; i++ {
				x, y, z = robotMove(repeatCmd, x, y, z)
			}
			repeat = 0
			repeatCmd = ""
		case repeat > 0 && s != '(' && s != ')':
			repeatCmd = repeatCmd + string(s)
		case s == 'L':
			z = (z + 1) % 4
		case s == 'R':
			z = (z - 1 + 4) % 4
		case s == 'F':
			switch {
			case z == Left || z == Right:
				x = x - z + 1
			case z == Top || z == Bottom:
				y = y - z + 2
			}
		case s == 'B':
			switch {
			case z == Left || z == Right:
				x = x + z - 1
			case z == Top || z == Bottom:
				y = y + z - 2
			}
		}

	}
	return
}
