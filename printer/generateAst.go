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

// capitalize makes the first letter uppercase
func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func defineVisitor(file *os.File, baseName string, types []string) {
	fmt.Fprintf(file, "type %sVisitor interface {\n", baseName)
	for _, t := range types {
		typeName := strings.TrimSpace(strings.Split(t, ":")[0])
		// Example: VisitBinaryExpr(expr Binary) interface{}
		fmt.Fprintf(file, "    Visit%s%s(%s %s) interface{}\n",
			typeName, baseName, strings.ToLower(baseName), typeName)
	}
	fmt.Fprintln(file, "}")
	fmt.Fprintln(file)
}

func defineType(file *os.File, baseName, className, fieldList string) {
	// Struct
	fmt.Fprintf(file, "type %s struct {\n", className)

	fields := strings.Split(fieldList, ", ")
	for _, field := range fields {
		parts := strings.Fields(field)
		if len(parts) >= 2 {
			fieldType := parts[0]
			fieldName := parts[1]

			if fieldType == "Object" {
				fieldType = "interface{}"
			}
			fmt.Fprintf(file, "    %s %s\n", capitalize(fieldName), fieldType)
		}
	}
	fmt.Fprintln(file, "}")
	fmt.Fprintln(file)

	// Accept method
	fmt.Fprintf(file, "func (n %s) Accept(visitor %sVisitor) interface{} {\n", className, baseName)
	fmt.Fprintf(file, "    return visitor.Visit%s%s(n)\n", className, baseName)
	fmt.Fprintln(file, "}")
	fmt.Fprintln(file)
}

func defineAst(outDir string, baseName string, types []string, extraImports []string) {
	path := filepath.Join(outDir, strings.ToLower(baseName)+".go")

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintln(file, "package ast")
	fmt.Fprintln(file)

	if len(extraImports) > 0 {
		fmt.Fprintln(file, "import (")
		for _, imp := range extraImports {
			fmt.Fprintf(file, "    %q\n", imp)
		}
		fmt.Fprintln(file, ")")
		fmt.Fprintln(file)
	}

	// Visitor
	defineVisitor(file, baseName, types)

	// Base interface
	fmt.Fprintf(file, "type %s interface {\n", baseName)
	fmt.Fprintf(file, "    Accept(visitor %sVisitor) interface{}\n", baseName)
	fmt.Fprintln(file, "}")
	fmt.Fprintln(file)

	// Concrete types
	for _, t := range types {
		parts := strings.Split(t, ":")
		className := strings.TrimSpace(parts[0])
		fieldList := strings.TrimSpace(parts[1])
		defineType(file, baseName, className, fieldList)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <output directory>")
		os.Exit(64)
	}

	outputDir := os.Args[1]

	defineAst(outputDir, "Expr", []string{
		"Binary   : Expr left, token.Token operator, Expr right",
		"Grouping : Expr expression",
		"Literal  : interface{} value",
		"Logical  : Expr left, token.Token operator, Expr right",
		"Unary    : token.Token operator, Expr right",
		"Variable : token.Token name",
		"Assign   : token.Token name, Expr value",
	}, []string{"github.com/shubhdevelop/YAPL/Token"})

	defineAst(outputDir, "Stmt", []string{
		"BlockStmt      : []Stmt statement",
		"ExpressionStmt : Expr expression",
		"IfStmt : Expr condition, Stmt thenBranch," +
			" Stmt elseBranch",
		"PrintStmt      : Expr expression",
		"VarStmt : token.Token name, Expr initializer",
		"WhileStmt: Expr condition, Stmt body",
		"BreakStmt: ",
	}, []string{"github.com/shubhdevelop/YAPL/Token"})
}
