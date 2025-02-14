package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jaideep-kalagara/minasm/util"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path")
		os.Exit(1)
	}

	// Ensure FileHandler and Stack are properly initialized
	fh := util.FileHandler{}
	stack := util.Stack{}

	program := fh.SplitFile(os.Args[1])
	for _, v := range program {
		parts := strings.Fields(v)
		if len(parts) == 0 {
			continue
		}

		opcode := parts[0]
		args := parts[1:]

		switch opcode {
		case "ADD":
			a, ok1 := stack.Pop()
			b, ok2 := stack.Pop()
			if !ok1 || !ok2 {
				fmt.Println("Not enough operands for ADD")
				os.Exit(1)
			}
			stack.Push(a + b)

		case "PUSH":
			if len(args) == 0 {
				fmt.Println("Missing argument for PUSH")
				os.Exit(1)
			}
			val, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid PUSH value:", args[0])
				os.Exit(1)
			}
			stack.Push(val)

		case "SUB":
			a, ok1 := stack.Pop()
			b, ok2 := stack.Pop()
			if !ok1 || !ok2 {
				fmt.Println("Not enough operands for SUB")
				os.Exit(1)
			}
			stack.Push(b - a)

		case "POP":
			val, ok := stack.Pop()
			if !ok {
				fmt.Println("Stack is empty")
				os.Exit(1)
			}
			fmt.Println(val)
		case "POPSTR":
			val, _ := strconv.Atoi(args[0])
			str := stack.PopString(val)
			fmt.Println(str)
		case "PRINT":
			if strings.Contains(v, "\"") {
				parts := strings.SplitN(v, "\"", 2)
				if len(parts) > 1 {
					fmt.Println(parts[1])
				} else {
					fmt.Println("Invalid PRINT statement")
				}
			} else {
				fmt.Println("Invalid PRINT format")
			}

		case "EXIT":
			os.Exit(0)

		case "READ":
			var input string
			fmt.Scan(&input) // Use &input to correctly read user input
			stack.PushString(input)

		default:
			fmt.Printf("Unknown opcode: %s\n", opcode)
			os.Exit(1)
		}
	}
}
