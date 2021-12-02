<a href="compiler.go"><img width="731" alt="THE SUPER TINY COMPILER" src="https://cloud.githubusercontent.com/assets/952783/14413766/134c4068-ff39-11e5-996e-9452973299c2.png"/></a>

You may have recently come across [The Super Tiny Compiler][1] project. I read
the code a handful of times back-to-back and decided to see if I could write a
compiler myself, in Go. This project essentially just a port of the original
one, which is written in JavaScript. With comments, it's around 1000 lines -
without, only around 250.

### Usage

```
$ git clone git@github.com:hazbo/the-super-tiny-compiler.git
$ cd the-super-tiny-compiler && go build -o tiny
$ ./tiny
```

Input: `(add 2 (subtract 10 5))`

Output: `add(2, subtract(10, 5));`

You can find the input towards the bottom of compiler.go.

### Tests

```
$ go test
```

### Contributing

Feel free to issue a pull request if you feel there is something you could or
would like to add to this project.

---

[![cc-by-4.0](https://licensebuttons.net/l/by/4.0/80x15.png)](http://creativecommons.org/licenses/by/4.0/)

[1]: https://github.com/thejameskyle/the-super-tiny-compiler
