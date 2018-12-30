# Monkey Programming Language and Interpreter

Following Thorsten Ball's Book [*Writing An Interpreter In Go*](https://interpreterbook.com/) 

This project is focused on parsing and evaluating a language called *Monkey*

## Monkey has the following features

- C-like syntax
- variable bindings
- integers and booleans
- arithmetic expressions
- built-in functions
- first-class and higher-order functions
- closures
- a string data structure
- an array data structure
- a hash data structure

## The Monkey interpreter built will have 

- Lexer
- Parser
- Abstract Syntax Tree (AST)
- Internal object system
- Evaluator

### Running Requirements 
- GO version above 1.0
- Change GOPATH to the project directory
  - [direnv](https://direnv.net/) works great for this and included in the project is a .envrc file that correctly reroutes the GOPATH to that project
  - On Windows you will need to [change the enviorment variable](https://github.com/golang/go/wiki/SettingGOPATH)

