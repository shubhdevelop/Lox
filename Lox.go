package main 

import (
	"fmt"
	"bufio"
	"os"
	"errors"
)

func run(source string){
	fmt.Println("executed: ", source)
}



func runFile(path string){
	fmt.Println("Running with the file:", path)
	bytes, err := os.ReadFile(path)
	if err != nil {
		errors.New("Error loading the file check the file path")
	}
	source := string(bytes[:]);

	run(source)
	fmt.Print(source);
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
