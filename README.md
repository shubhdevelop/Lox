# YAPL Interpreter

A Go implementation of a complete programming language interpreter

## Overview

YAPL (Yet Another Programming Language) is a dynamically-typed, interpreted programming language with a clean and simple syntax. This implementation provides a full-featured interpreter with lexical analysis, parsing, and execution capabilities, including control flow statements, variable scoping, and comprehensive error handling.

## Language Syntax

### Basic Syntax Rules

1. **Statements** must end with a semicolon (`;`)
2. **Identifiers** can contain letters, digits, and underscores, but cannot start with a digit
3. **Keywords** are reserved and cannot be used as identifiers
4. **Case sensitivity**: Lox is case-sensitive
5. **Whitespace** is generally ignored except for separating tokens

### Reserved Keywords

```
and, class, else, false, for, fun, if, nil, or, print, return, super, this, true, var, while
```

**Note**: `class`, `fun`, `return`, `super`, and `this` are reserved for future implementation.

### Operators

| Operator | Description | Example |
|----------|-------------|---------|
| `+` | Addition / String concatenation | `1 + 2` or `"hello" + "world"` |
| `-` | Subtraction / Negation | `5 - 3` or `-x` |
| `*` | Multiplication | `3 * 4` |
| `/` | Division | `10 / 2` |
| `!` | Logical NOT | `!true` |
| `==` | Equality | `x == y` |
| `!=` | Inequality | `x != y` |
| `>` | Greater than | `5 > 3` |
| `>=` | Greater than or equal | `5 >= 5` |
| `<` | Less than | `3 < 5` |
| `<=` | Less than or equal | `3 <= 5` |
| `and` | Logical AND | `true and false` |
| `or` | Logical OR | `true or false` |

## Usage

### Building the Interpreter

```bash
go build
```

### Running Programs

#### From a File
```bash
./Lox script.yapl
```

#### Interactive Mode
```bash
./Lox
```

In interactive mode, you can:
- Enter expressions and statements
- Type `clear` to clear the screen
- Type `exit` to quit

### Example Programs

#### Basic Arithmetic
```yapl
var a = 10;
var b = 20;
print a + b;  // Output: 30
print a * b;  // Output: 200
```

#### String Operations
```yapl
var greeting = "Hello";
var name = "World";
print greeting + " " + name;  // Output: Hello World
```

#### Boolean Logic
```yapl
var x = true;
var y = false;
print y == x;  // Output: false
print !x;      // Output: false
print true and false;  // Output: false
print true or false;   // Output: true
```

#### Variable Assignment
```yapl
var count = 0;
print count;    // Output: 0
count = count + 1;
print count;    // Output: 1
```

#### Control Flow
```yapl
// If statement
var age = 18;
if (age >= 18) {
    print "You are an adult";
} else {
    print "You are a minor";
}

// While loop
var i = 0;
while (i < 5) {
    print i;
    i = i + 1;
}

// For loop (desugared to while)
for (var j = 0; j < 3; j = j + 1) {
    print "Iteration: " + j;
}
```

#### Block Scoping
```yapl
var global = "I'm global";
{
    var local = "I'm local";
    print global;  // Output: I'm global
    print local;   // Output: I'm local
}
print global;  // Output: I'm global
// print local;  // Error: Undefined variable 'local'
```

## Project Structure

```
YAPL/
├── Token/           # Token definitions and types
├── Scanner/         # Lexical analysis (tokenization)
├── parser/          # Syntax analysis (parsing)
├── ast/             # Abstract Syntax Tree nodes
├── Interpreter/     # Expression and statement evaluation
├── environment/     # Variable environment management with scoping
├── YaplErrors/      # Error handling and reporting
├── state/           # Global interpreter state
├── printer/         # AST pretty printing utilities
├── main.go          # Main interpreter entry point
├── test.lox         # Example YAPL program
└── main.yapl        # Additional example program
```

## Architecture

The interpreter follows a traditional pipeline architecture:

1. **Scanner**: Converts source code into tokens
2. **Parser**: Builds an Abstract Syntax Tree (AST) from tokens
3. **Interpreter**: Evaluates the AST and produces results

### Key Components

- **Token**: Represents lexical units (keywords, operators, literals)
- **Scanner**: Implements lexical analysis with support for comments, strings, numbers, and identifiers
- **Parser**: Recursive descent parser with error recovery and support for all control flow statements
- **AST**: Tree representation of program structure with expression and statement nodes
- **Interpreter**: Visitor pattern implementation for expression and statement evaluation
- **Environment**: Manages variable storage and lookup with proper scoping support
- **Error Handling**: Comprehensive error reporting for lexical, parse, and runtime errors

## Error Handling

The interpreter provides comprehensive error reporting:

- **Lexical Errors**: Invalid characters, unterminated strings
- **Parse Errors**: Syntax errors with line numbers and helpful messages
- **Runtime Errors**: Type mismatches, undefined variables

Error messages include:
- Line number where the error occurred
- Description of the error
- Context around the error location

## Development Status

This implementation represents a mature interpreter with comprehensive language features. The core functionality is complete and includes:

### ✅ Implemented Features

- **Complete Expression System**: All arithmetic, comparison, logical, and unary operations
- **Control Flow**: `if`/`else` statements, `while` loops, and `for` loops (desugared to while)
- **Variable Management**: Declaration, assignment, and proper scoping with block environments
- **Data Types**: Numbers (float64), strings, booleans, and nil
- **Error Handling**: Comprehensive lexical, parse, and runtime error reporting
- **Interactive Mode**: REPL with clear and exit commands
- **File Execution**: Run YAPL programs from files

### 🚧 Future Enhancements

1. **Functions**: User-defined functions with parameters and return values
2. **Classes and Objects**: Object-oriented programming support
3. **Standard Library**: Built-in functions for common operations
4. **Modules**: Import/export system for code organization
5. **Advanced Error Recovery**: Better error messages and suggestions


## Features

#### **Data Types**
- **Numbers**: Floating-point numbers, all numbers are double, with 2-place float precision (e.g., `42`, `3.14`) 
- **Strings**: Text literals enclosed in double quotes (e.g., `"hello world"`)
- **Booleans**: `true` and `false`
- **Nil**: Represents the absence of a value

#### **Expressions**
- **Arithmetic Operations**:
  - Addition: `+` (supports numbers and string concatenation)
  - Subtraction: `-`
  - Multiplication: `*`
  - Division: `/`
- **Comparison Operations**:
  - Equality: `==`, `!=`
  - Relational: `>`, `>=`, `<`, `<=`
- **Logical Operations**:
  - Logical AND: `and`
  - Logical OR: `or`
  - Logical NOT: `!`
- **Unary Operations**:
  - Negation: `-` (for numbers)
  - Logical NOT: `!`
- **Grouping**: Parentheses `()` for expression precedence
- **Variable Access**: Direct variable name references

#### **Statements**
- **Variable Declaration**: `var name = value;`
- **Print Statement**: `print expression;`
- **Expression Statement**: Any expression followed by `;`
- **If Statement**: `if (condition) statement else statement`
- **While Loop**: `while (condition) statement`
- **For Loop**: `for (initializer; condition; increment) statement`
- **Block Statement**: `{ statement1; statement2; ... }`

#### **Variables**
- **Declaration**: `var variableName;` or `var variableName = initialValue;`
- **Assignment**: `variableName = newValue;`
- **Scope**: Block-scoped variables with proper environment nesting
- **Lookup**: Variables are looked up in the current scope and enclosing scopes

#### **Comments**
- **Single-line comments**: `// This is a comment`

#### **Error Handling**
- **Lexical Errors**: Invalid characters, unterminated strings
- **Parse Errors**: Syntax errors with helpful error messages
- **Runtime Errors**: Type errors, undefined variables
