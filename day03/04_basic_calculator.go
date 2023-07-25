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

type Stack[T comparable] struct {
	array []T
}

func (stack *Stack[T]) Push(x T) {
	stack.array = append(stack.array, x)
}

func (stack *Stack[T]) Pop() T {
	val := stack.array[len(stack.array) - 1]
	stack.array = stack.array[ :len(stack.array) - 1]
	return val
}

func (stack *Stack[T]) Size() int {
	return len(stack.array)
}

func (stack *Stack[T]) Peek() T {
	return stack.array[len(stack.array) - 1]
}

type Calculator struct {
	expression string
	operations *Stack[rune]
	operands *Stack[int]
	precendence Array[rune]
}

func (Calc *Calculator) execute() int {
	operand := 0

	for _, c := range Calc.expression {
		if unicode.IsDigit(c) {
			operand = operand * 10 + (int(c) - '0')
		} else if Calc.precendence.Contains(c) {
			Calc.operands.Push(operand)
			operand = 0
			if Calc.operations.Size() > 0 {
				top := Calc.operations.Peek()
				for Calc.operations.Size() > 0  && Calc.precendence.indexOf(c) <= Calc.precendence.indexOf(top) {
					Calc.operations.Pop()
					Operand2 := Calc.operands.Pop()
					Operand1 := Calc.operands.Pop()
					if top == '*' {
						Calc.operands.Push(Operand1 * Operand2)
					} else {
						Calc.operands.Push(Operand1 + Operand2)
					}
					if Calc.operations.Size() > 0  {
						top = Calc.operations.Peek()
					}
				}
			}
			
			Calc.operations.Push(c)
		}
	}

	result := Calc.operands.Pop()
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
			operations : &Stack[rune] {
				array : make([]rune, 0),
			},
			operands : &Stack[int] {
				array : make([]int, 0),
			},
		}
	
		calc.precendence = Array[rune] { '#', '+', '*', }
		fmt.Print("result is :   ")
		fmt.Println(calc.execute())
		fmt.Println(strings.Repeat("_", 20))
		fmt.Println()
	}
}
