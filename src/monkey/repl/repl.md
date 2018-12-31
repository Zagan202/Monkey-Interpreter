# REPL

An interactive prompt is also called a “REPL” (pronounced like “rebel” but with a “p”). The name comes from Lisp where implementing one is as simple as wrapping a loop around a few built-in functions:

`(print (eval (read)))`
Working outwards from the most nested call, you Read a line of input, Evaluate it, Print the result, then Loop and do it all over again.[^first].



[^first]: [Crafting Interpreters](http://www.craftinginterpreters.com/scanning.html)

