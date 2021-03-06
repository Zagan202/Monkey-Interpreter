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

### MONKEY INTERPRETER TODO

- [x] Lexer, REPL, Token
- [X] Lexer Write up
- [X] REPL Write up
- [X] Token Write up
- [X] Lexer,REPL, Token Code Comments
- [ ] Parser
- [ ] Parser Code Comments
- [ ] Parser Write up
- [ ] Eval
- [ ] Eval Code Comments
- [ ] Eval Write up


### Running Requirements 
- Makefile included!
  - make
    - run main code
  - make clean 
    - clean any binaries created 
- GO version above 1.0
- Change GOPATH to the project directory
  - [direnv](https://direnv.net/) works great for this and included in the project is a .envrc file that correctly reroutes the GOPATH to that project
  - On Windows you will need to [change the enviorment variable](https://github.com/golang/go/wiki/SettingGOPATH)

