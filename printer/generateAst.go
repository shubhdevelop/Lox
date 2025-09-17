//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// capitalize capitalizes the first letter of a string
func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func defineVisitor(file *os.File, baseName string, types []string) {
	fmt.Fprintln(file, "type Visitor interface {")
	for _, t := range types {
		typeName := strings.TrimSpace(strings.Split(t, ":")[0])
		// Fixed: Generate proper method signature with parameter name and type
		fmt.Fprintf(file, "    Visit%s%s(expr %s) interface{}\n",
			typeName, baseName, typeName)
	}
	fmt.Fprintln(file, "}")
	fmt.Fprintln(file)
}

func defineType(file *os.File, baseName, className, fieldList string) {
	// Struct definition
	fmt.Fprintf(file, "type %s struct {\n", className)

	// Fields
	fields := strings.Split(fieldList, ", ")
	for _, field := range fields {
		field = strings.TrimSpace(field)
		parts := strings.Fields(field) // Use Fields instead of Split to handle multiple spaces
		if len(parts) >= 2 {
			// Fixed: Correct order - first part is type, second is name
			fieldType := parts[0]
			fieldName := parts[1]

			// Map Java "Object" to Go "interface{}"
			if fieldType == "Object" {
				fieldType = "interface{}"
			}

			// Fixed: Use proper capitalization
			fmt.Fprintf(file, "    %s %s\n", capitalize(fieldName), fieldType)
		}
	}
	fmt.Fprintln(file, "}")
	fmt.Fprintln(file)

	// Accept method
	fmt.Fprintf(file, "func (e %s) Accept(visitor Visitor) interface{} {\n", className)
	fmt.Fprintf(file, "    return visitor.Visit%s%s(e)\n", className, baseName)
	fmt.Fprintln(file, "}")
	fmt.Fprintln(file)
}

func defineAst(outDir string, baseName string, types []string) {
	path := filepath.Join(outDir, strings.ToLower(baseName)+".go")

	// Create file
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write package
	fmt.Fprintln(file, "package ast")
	fmt.Fprintln(file)

	// Imports (if needed)
	fmt.Fprintln(file, "import (")
	fmt.Fprintln(file, `    "github.com/shubhdevelop/Lox/token"`)
	fmt.Fprintln(file, ")")
	fmt.Fprintln(file)

	// Define visitor interface first
	defineVisitor(file, baseName, types)

	// Define base interface
	fmt.Fprintf(file, "type %s interface {\n", baseName)
	fmt.Fprintln(file, "    Accept(visitor Visitor) interface{}")
	fmt.Fprintln(file, "}")
	fmt.Fprintln(file)

	// Define concrete types
	for _, t := range types {
		parts := strings.Split(t, ":")
		if len(parts) >= 2 {
			className := strings.TrimSpace(parts[0])
			fieldList := strings.TrimSpace(parts[1])
			defineType(file, baseName, className, fieldList)
		}
	}
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: go run main.go <output directory>")
		os.Exit(64) // like in Crafting Interpreters
	}

	outputDirectory := args[1]
	defineAst(outputDirectory, "Expr", []string{
		"Binary   : Expr left, token.Token operator, Expr right",
		"Grouping : Expr expression",
		"Literal  : interface{} value",
		"Unary    : token.Token operator, Expr right",
	})
}
