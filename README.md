# The Super Tiny Compiler

You may have recently come across [The Super Tiny Compiler][1] project. I read
the code a handful of times back-to-back and decided to see if I could write a
compiler myself, in Go. This project is *heavily* inspired from the original
repo that is written in JavaScript.

It's still pretty rough around the edges, and I will be improving it over time.
But it works! :)

### Usage

```
$ git clone git@github.com:hazbo/the-super-tiny-compiler.git
$ cd the-super-tiny-compiler && go build -o tiny
$ ./tiny
```

Input: `(add 2 (subtract 10 5))`

Output: `add(2, subtract(10, 5));`

You can find the input towards the bottom of compiler.go.

P.S. Sorry in advance for the (or lack of) error handling. I will come around to
this. Unless you want to do it. Also, comments similar to the comments in the
original project will be added.

[![cc-by-4.0](https://licensebuttons.net/l/by/4.0/80x15.png)](http://creativecommons.org/licenses/by/4.0/)

[1]: https://github.com/thejameskyle/the-super-tiny-compiler
