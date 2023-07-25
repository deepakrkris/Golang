package main

import (
	"fmt"
	"unicode"
	"strings"
)	

type Array[T comparable] []T

func (arr Array[T]) Contains (x T) bool {
    for _, v := range arr {
        if v == x {
            return true
        }
    }
    return false
}

func (arr Array[T]) indexOf (x T) int {
    for i, v := range arr {
        if v == x {
            return i
        }
    }
    return -1
}

type Calculator struct {
	expression string
	operations Array[rune]
	operands []int
	precendence Array[rune]
}

func (Calc *Calculator) execute() int {
	operand := 0

	for _, c := range Calc.expression {
		if unicode.IsDigit(c) {
			operand = operand * 10 + (int(c) - '0')
		} else if Calc.precendence.Contains(c) {
			//fmt.Println(operand)
			Calc.operands = append(Calc.operands, operand)
			operand = 0
			topIndex := len(Calc.operations) - 1
			if topIndex >= 0 {
				top := Calc.operations[topIndex]
				for len(Calc.operations) > 0 && Calc.precendence.indexOf(c) <= Calc.precendence.indexOf(top) {
					Calc.operations = Calc.operations[:topIndex]
					Operand2 := Calc.operands[len(Calc.operands) - 1]
					Operand1 := Calc.operands[len(Calc.operands) - 2]
					Calc.operands = Calc.operands[ : len(Calc.operands) - 2]
					if top == '*' {
						Calc.operands = append(Calc.operands, Operand1 * Operand2)
					} else {
						Calc.operands = append(Calc.operands, Operand1 + Operand2)
					}
					if len(Calc.operations) > 0  {
						topIndex = len(Calc.operations) - 1
						top = Calc.operations[topIndex]
					}
				}
			}
			Calc.operations = append(Calc.operations, c)
		}
	}

	topIndex := len(Calc.operands) - 1
	result := Calc.operands[topIndex]
	return result
}

func test_data() []string {
	return []string {
		"1+2#",
		"2+2*3#",
		"2+2*3*2+1#",
		"2+22*2#",
	}
}

func main() {

	for _, data := range test_data() {
		fmt.Println()
		fmt.Println(strings.Repeat("_", 20))
		fmt.Println(" data is  ", data)

		calc := Calculator {
			expression : data,
			operations : make(Array[rune], 0),
			operands : make([]int, 0),
		}
	
		calc.precendence = Array[rune] { '#', '+', '*', }
		fmt.Print("result is :   ")
		fmt.Println(calc.execute())
		fmt.Println(strings.Repeat("_", 20))
		fmt.Println()
	}
}
