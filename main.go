package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/jaideep-kalagara/minasm/util"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Please provide a file path")
		os.Exit(1)
	}

	var fh util.FileHandler

	var stack util.Stack
	program := fh.SplitFile(os.Args[1])
	for i, v := range program {
		opcode := strings.Split(v, " ")[0]
		args := strings.Split(v, " ")[1:]
		fmt.Println(i, opcode, args)
		switch opcode {
		case "ADD":
			a, _ := stack.Pop()
			b, ok := stack.Pop()
			if !ok {
				fmt.Println("Not enough operands")
				os.Exit(1)
			}
			stack.Push(a + b)
		case "PUSH":
			val, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid value")
				os.Exit(1)
			}
			stack.Push(val)
		case "SUB":
			a, _ := stack.Pop()
			b, ok := stack.Pop()
			if !ok {
				fmt.Println("Not enough operands")
				os.Exit(1)
			}
			stack.Push(b - a)
		case "POP":
			if len(stack.Stack) == 0 {
				fmt.Println("Stack is empty")
				os.Exit(1)
			}
			val, ok := stack.Pop()
			if !ok {
				fmt.Println("Error popping from stack")
				os.Exit(1)
			}
			fmt.Println(val)
		}
	}

}
