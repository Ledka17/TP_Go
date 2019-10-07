package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		i := calc(os.Args[1])
		fmt.Println(i)
	}
}

func calc(expr string) int {
	return eval(toPpn(expr))
}

func toPpn(expr string) []string {
	var rpn []string
	var buf []string
	expr = backspace(expr)
	exprArray := strings.Split(expr, " ") // для разделителя пробел
	buf = append(buf, "|") // символ начала выражения
	for i := 0; i < len(exprArray); i++ {
		prev := buf[len(buf) - 1]
		cur := exprArray[i]
		if isNum(cur) {
			rpn = append(rpn, cur)
		} else if ((isSum(cur) || cur ==")") && (prev != "(" && prev != "|")) || (isFactor(cur) && isFactor(prev)) { // 2
			rpn = append(rpn, prev)
			buf = buf[:len(buf) - 1]
			i--
		} else if prev == "(" && cur == ")" { // 3
			buf = buf[:len(buf) - 1]
		} else { // 1
			buf = append(buf, cur)
		}
	}
	for ; len(buf) > 1; {
		rpn = append(rpn, buf[len(buf) - 1])
		buf = buf[:len(buf) - 1]
	}
	return rpn
}

func eval(expr []string) int {
	var buf []string
	for i := 0; i < len(expr); i++ {
		if isNum(expr[i]) {
			buf = append(buf, expr[i])
		} else {
			if len(buf) > 1 {
				num1 := buf[len(buf) - 2]
				num2 := buf[len(buf) - 1]
				buf = buf[:len(buf) - 2]
				buf = append(buf, operate(num1, num2, expr[i]))
			}
		}
	}
	return atoi(buf[0])
}

func atoi(element string) int {
	elementInt, _ := strconv.Atoi(element)
	return elementInt
}

func backspace(expr string) string{
	for i := len(expr) - 2; i > 0; i-- {
		if isFactor(string(expr[i])) || isSum(string(expr[i])) {
			expr = expr[:i] + " " + string(expr[i]) + " " + expr[i + 1:]
			i--
		}
	}
	return expr
}

func operate(num1 string, num2 string, op string) string {
	switch op {
	case "+":
		return strconv.Itoa(atoi(num1) + atoi(num2))
	case "-":
		return strconv.Itoa(atoi(num1) - atoi(num2))
	case "/":
		return strconv.Itoa(atoi(num1) / atoi(num2))
	case "*":
		return strconv.Itoa(atoi(num1) * atoi(num2))
	}
	return num1
}

func isSum(expr string) bool {
	if expr == "+" || expr == "-" {
		return true
	}
	return false
}

func isFactor(expr string) bool {
	if expr == "*" || expr == "/" {
		return true
	}
	return false
}

func isNum(expr string) bool {
	if expr != "+" && expr != "-" && expr != "/" && expr != "*" {
		return true
	}
	return false
}