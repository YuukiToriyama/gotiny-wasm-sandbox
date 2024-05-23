package main

import "strconv"

func main() {}

//export console_log
func log(message string)

//export alert
func alert(message string)

//export multiply
func multiply(left, right int) int {
	log("left: " + strconv.Itoa(left) + "\n" + "right: " + strconv.Itoa(right))
	alert("掛け算を実行します！")
	return left * right
}
