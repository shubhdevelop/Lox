package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"github.com/shubhdevelop/Lox/scanner"
	"github.com/shubhdevelop/Lox/state"
)

func run(source string){
	scanner := scanner.Scanner{Source:source}
	tokens, err := scanner.ScanToken()
	if err != nil {
		errors.New("Error Scanning tokens")
	} else {
		fmt.Println(tokens[:])
	}
}


func runFile(path string){
	fmt.Println("Running with the file:", path)
	bytes, err := os.ReadFile(path)
	if err != nil {
		errors.New("Error loading the file check the file path")
	}
	source := string(bytes[:]);

	run(source)
	if state.HadError { 
		os.Exit(65)
	}

}   

func runPrompt(){
	fmt.Println("Running in prompt Mode")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")
		line,err := reader.ReadString('\n')
		if err != nil {
			errors.New("Error Reading the line")
			break;
		} else if len(line) == 0 {    
			break
		} else if line == "exit()\n"{
			break
		}
		run(string(line));
		if state.HadError { 
			os.Exit(65)
		}
	}

	fmt.Println("exiting out of LOX")

}

func main(){
	args := os.Args[1:]

	if len(args) > 1 {
		errors.New("usage Lox [script]")
	} else if len(args) == 1{
		runFile(args[0])
	}else{
		runPrompt()
	}

	fmt.Println("welcome to lox world")
}
