package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	interpreter "github.com/shubhdevelop/Lox/Interpreter"
	"github.com/shubhdevelop/Lox/parser"
	"github.com/shubhdevelop/Lox/scanner"
	"github.com/shubhdevelop/Lox/state"
)

func run(source string) {
	state.HadError = false // Reset error state
	scanner := scanner.Scanner{Source: source}
	tokens, err := scanner.ScanTokens()
	interpreter := interpreter.Interpreter{}
	parserInstance := parser.Parser{
		Tokens: tokens,
	}
	if err != nil {
		fmt.Println(errors.New("Error Scanning tokens"))
		state.HadError = true
		return
	}
	expr := parserInstance.Parse()
	interpreter.Interpret(expr)
	if state.HadError {
		return
	}
}

func runFile(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(errors.New("Error loading the file check the file path"))
	}
	source := string(bytes[:])

	run(source)
	if state.HadError {
		os.Exit(65)
	}
	if state.HadRuntimeError {
		os.Exit(70)
	}

}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error Reading the line")
			continue
		} else if line == "\n" || line == "" {
			continue
		} else if line == "clear\n" {
			fmt.Print("\033[H\033[2J")
			continue
		} else if line == "exit\n" {
			break
		}
		run(string(line))
		if state.HadError {
			state.HadError = false // Reset error state for next input
		}
	}
}

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		panic(errors.New("usage Lox [script]"))
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}
