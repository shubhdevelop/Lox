# YAPL Interpreter

A Go implementation of a programming language interpreter

## Overview

YAPL is a dynamically-typed, interpreted programming language with a clean and simple syntax. This implementation provides a complete interpreter with lexical analysis, parsing, and execution capabilities.

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
| `and` | Logical AND | `true and false` | (yet to be implemented)
| `or` | Logical OR | `true or false` | (yet to be implemented)

## Usage

### Building the Interpreter

```bash
go build
```

### Running Programs

#### From a File
```bash
./Lox script.lox
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
```lox
var a = 10;
var b = 20;
print a + b;  // Output: 30
print a * b;  // Output: 200
```

#### String Operations
```lox
var greeting = "Hello";
var name = "World";
print greeting + " " + name;  // Output: Hello World
```

#### Boolean Logic
```lox
var x = true;
var y = false;
print y == x;
print !x;       // Output: false
```

#### Variable Assignment
```lox
var count = 0;
print count;    // Output: 0
count = count + 1;
print count;    // Output: 1
```

## Project Structure

```
Lox/
├── Token/           # Token definitions and types
├── Scanner/         # Lexical analysis (tokenization)
├── parser/          # Syntax analysis (parsing)
├── ast/             # Abstract Syntax Tree nodes
├── Interpreter/     # Expression and statement evaluation
├── environment/     # Variable environment management
├── YaplErrors/       # Error handling and reporting
├── state/           # Global interpreter state
├── printer/         # AST pretty printing utilities
└── main.yapl        # Example Lox program
```

## Architecture

The interpreter follows a traditional pipeline architecture:

1. **Scanner**: Converts source code into tokens
2. **Parser**: Builds an Abstract Syntax Tree (AST) from tokens
3. **Interpreter**: Evaluates the AST and produces results

### Key Components

- **Token**: Represents lexical units (keywords, operators, literals)
- **Scanner**: Implements lexical analysis with support for comments, strings, numbers, and identifiers
- **Parser**: Recursive descent parser with error recovery
- **AST**: Tree representation of program structure
- **Interpreter**: Visitor pattern implementation for expression and statement evaluation
- **Environment**: Manages variable storage and lookup

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

This implementation represents the early stages of a complete interpreter. The core expression evaluation and basic statements are functional, providing a solid foundation for adding more advanced language features.

### Next Steps

1. Implement control flow statements (`if`, `while`, `for`)
2. Add function support
3. Implement classes and objects
4. Add local scoping
5. Enhance error handling and debugging features


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

#### **Variables**
- **Declaration**: `var variableName;` or `var variableName = initialValue;`
- **Assignment**: `variableName = newValue;`
- **Scope**: Global scope (variables are accessible throughout the program)

#### **Comments**
- **Single-line comments**: `// This is a comment`

#### **Error Handling**
- **Lexical Errors**: Invalid characters, unterminated strings
- **Parse Errors**: Syntax errors with helpful error messages
- **Runtime Errors**: Type errors, undefined variables
