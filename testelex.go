package main

import (
	"bufio"
	"fmt"
	"os"
)


/*
 * Complete the 'toString' function below.
 *
 * The function is expected to return a STRING.
 */

type Node struct {
	Value any
	Left Node
	Right Node
}


const mapOperations = map[string]func{ "+": func (a, b int) { return a + b },
	"*": func (a, b int) { return a * b },
	"/": func (a, b int) { return a / b },
	"sqrt": func (a, b int) { return a / b },}


func eval(a Node) (int, error) {
	if a.Left == nil && a.Right == nil {
		return a.value
	}

	resultLeft, err := eval(a.Left)
	resultRight, err := eval(a.Right)

	cb, ok := mapOperations[a.value]
	if !ok {
		return 0,
	}
	return cb(resultLeft, resultRight), nil
}

const opS = []


func toString(a Node) (string, bool) {
	needP := false

	if a.Left == nil && a.Right == nil {
		return a.value, needP
	}

	if a.value == "+" || a.value == "-" {
		needP = true
	}

	if a.value == "*" {
		needP = true
	}

	resultLeft, leftHasChildren := eval(a.Left)
	resultRight, rightHasChildren := eval(a.Right)

	if leftHasChildren {
		resultLeft = "(" + resultLeft + ")"
	}

	if rightHasChildren {
		resultRight = "(" + resultRight + ")"
	}

	return resultLeft + a.value + resultRight, needP
}


func main() {

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	result := toString()

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
