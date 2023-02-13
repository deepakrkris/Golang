package main

import (
	"fmt"
	"unicode"
	"strings"
	"errors"
	"github.com/emirpasic/gods/stacks/arraystack"
)

type Calculator3 struct {
	expression string
	operations *arraystack.Stack
	operands *arraystack.Stack
	precendence []rune
}

func Contains2[T comparable] (arr [] T, x T) bool {
    for _, v := range arr {
        if v == x {
            return true
        }
    }
    return false
}

func indexOf2[T comparable] (arr [] T, x T) int {
    for i, v := range arr {
        if v == x {
            return i
        }
    }
    return -1
}

func (Calc *Calculator3) execute() (int, error) {
	operand := 0

	for _, c := range Calc.expression {
		if unicode.IsDigit(c) {
			operand = operand * 10 + (int(c) - '0')
		} else if Contains2(Calc.precendence, c) {
			//fmt.Println(operand)
			Calc.operands.Push(operand)
			operand = 0
			if Calc.operations.Size() > 0 {
				top , _ := Calc.operations.Peek()
				current_operator, ok := top.(rune)
				if !ok {
					return -1, errors.New("conversion error")
				}

				for Calc.operations.Size() > 0 && indexOf2(Calc.precendence, c) <= indexOf2(Calc.precendence, current_operator) {
					Operand2, _ := Calc.operands.Pop()
					Operand1, _ := Calc.operands.Pop()
					Calc.operations.Pop()

					value2, _ := Operand2.(int)
					value1, _ := Operand1.(int)

					if current_operator == '*' {
						//fmt.Println("*", value1, value2)
						Calc.operands.Push(value1 * value2)
					} else if current_operator == '+' {
						//fmt.Println("+", value1, value2)
						Calc.operands.Push(value1 + value2)
					} else if current_operator == '/' {
						//fmt.Println("/", value1, value2)
						Calc.operands.Push(value1 / value2)
					}
					if Calc.operations.Size() > 0 {
						top , _ = Calc.operations.Peek()
						current_operator, ok = top.(rune)
						if !ok {
							return -1, errors.New("conversion error")
						}
					}
				}
			}
			Calc.operations.Push(c)
		}
	}

	top, _ := Calc.operands.Peek()
	result, _ := top.(int)
	return result, nil
}

func __test_data() []string {
	return []string {
		"1+2#",
		"2+2*3#",
		"2+2*3*2+1#",
		"2+22/2*2#",
		"2+2*2*22/2+1#",
	}
}

func main() {

	for _, data := range __test_data() {
		fmt.Println()
		fmt.Println(strings.Repeat("_", 20))
		fmt.Println(" data is  ", data)

		calc := Calculator3 {
			expression : data,
		}
		calc.operands = arraystack.New()
		calc.operations = arraystack.New()
		calc.precendence = []rune { '#', '+', '*', '/'}
		fmt.Print("result is :   ")
		result, _ := calc.execute()
		fmt.Println(result)
		fmt.Println(strings.Repeat("_", 20))
		fmt.Println()
	}
}
