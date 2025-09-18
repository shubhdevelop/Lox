package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/shubhdevelop/Lox/parser"
	"github.com/shubhdevelop/Lox/printer"
	"github.com/shubhdevelop/Lox/scanner"
	"github.com/shubhdevelop/Lox/state"
	"os"
)

func run(source string) {
	scanner := scanner.Scanner{Source: source}
	tokens, err := scanner.ScanTokens()
	parserInstance := parser.Parser{
		Tokens: tokens,
	}
	if err != nil {
		fmt.Println(errors.New("Error Scanning tokens"))
	}
	expr := parserInstance.Parse()
	if state.HadError {
		return
	}
	printer := printer.AstPrinter{}
	fmt.Printf("%v \n", printer.Print(expr))
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

}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(errors.New("Error Reading the line"))
			continue
		} else if len(line) == 0 {
			break
		} else if line == "exit()\n" {
			break
		}
		run(string(line))
		if state.HadError {
			os.Exit(65)
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
