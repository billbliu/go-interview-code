/*
 * @Author: bill
 * @Date: 2022-01-15 17:02:01
 * @LastEditors: bill
 * @LastEditTime: 2022-01-15 17:03:36
 * @Description: go test -v 0006_robot_move.go 0006_robot_move_test.go
 * @FilePath: /go-interview-code/questions/0006_robot_move_test.go
 */
package questions

import (
	"log"
	"testing"
)

func TestRobotMove(t *testing.T) {
	x, y, z := robotMove("R2(LF)", 0, 0, Top)
	log.Printf("x: %d, y: %d, z: %d", x, y, z)
}
