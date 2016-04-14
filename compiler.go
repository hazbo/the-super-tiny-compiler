package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type token struct {
	kind  string
	value string
}

func tokenizer(input string) []token {
	input += "\n"

	current := 0
	tokens := []token{}

	for current < len([]rune(input)) {
		char := string([]rune(input)[current])

		if char == "(" {
			tokens = append(tokens, token{
				kind:  "paren",
				value: "(",
			})
			current++
			continue
		}
		if char == ")" {
			tokens = append(tokens, token{
				kind:  "paren",
				value: ")",
			})
			current++
			continue
		}
		if match(" ", char) == true {
			current++
			continue
		}
		if match("[0-9]", char) == true {
			value := ""

			for match("[0-9]", char) {
				value += char
				current++
				char = string([]rune(input)[current])
			}
			tokens = append(tokens, token{
				kind:  "number",
				value: value,
			})
			continue
		}
		if match("[a-z]", char) == true {
			value := ""
			for match("[a-z]", char) {
				value += char
				current++
				char = string([]rune(input)[current])
			}

			tokens = append(tokens, token{
				kind:  "name",
				value: value,
			})
			continue
		}
		break
	}
	return tokens
}

func match(pattern, char string) bool {
	m, err := regexp.MatchString(pattern, char)
	if err != nil {
		log.Fatal(err)
	}
	return m
}

type node struct {
	kind       string
	value      string
	name       string
	callee     *node
	expression *node
	body       []node
	params     []node
	arguments  *[]node
	context    *[]node
}

type nodes []node
type ast node

// parser counter
var pc int

// parser token
var pt []token

func parser(tokens []token) ast {
	pc = 0
	pt = tokens

	ast := ast{
		kind: "Program",
		body: []node{},
	}

	for pc < len(pt) {
		ast.body = append(ast.body, walk())
	}

	return ast
}

func walk() node {
	token := pt[pc]

	if token.kind == "number" {
		pc++
		return node{
			kind:  "NumberLiteral",
			value: token.value,
		}
	}
	if token.kind == "paren" && token.value == "(" {
		pc++
		token = pt[pc]

		n := node{
			kind:   "CallExpression",
			name:   token.value,
			params: []node{},
		}

		pc++
		token = pt[pc]

		for token.kind != "paren" || (token.kind == "paren" && token.value != ")") {
			n.params = append(n.params, walk())
			token = pt[pc]
		}

		pc++
		return n
	}

	log.Fatal("err")
	return node{}
}

type visitor map[string]func(n *node, p node)

func traverser(a ast, v visitor) {
	traverseNode(node(a), node{}, v)
}

func traverseArray(a []node, p node, v visitor) {
	for _, child := range a {
		traverseNode(child, p, v)
	}
}

func traverseNode(n, p node, v visitor) {
	for k, va := range v {
		if k == n.kind {
			va(&n, p)
		}
	}

	switch n.kind {
	case "Program":
		traverseArray(n.body, n, v)
		break
	case "CallExpression":
		traverseArray(n.params, n, v)
		break
	case "NumberLiteral":
		break
	default:
		log.Fatal("err")
	}
}

func transformer(a ast) ast {
	nast := ast{
		kind: "Program",
		body: []node{},
	}

	a.context = &nast.body

	traverser(a, map[string]func(n *node, p node){
		"NumberLiteral": func(n *node, p node) {
			*p.context = append(*p.context, node{
				kind:  "NumberLiteral",
				value: n.value,
			})
		},
		"CallExpression": func(n *node, p node) {
			e := node{
				kind: "CallExpression",
				callee: &node{
					kind: "Identifier",
					name: n.name,
				},
				arguments: new([]node),
			}

			n.context = e.arguments

			if p.kind != "CallExpression" {
				es := node{
					kind:       "ExpressionStatement",
					expression: &e,
				}
				*p.context = append(*p.context, es)
			} else {
				*p.context = append(*p.context, e)
			}

		},
	})
	return nast
}

func codeGenerator(n node) string {
	switch n.kind {
	case "Program":
		var r []string
		for _, no := range n.body {
			r = append(r, codeGenerator(no))
		}
		return strings.Join(r, "\n")

	case "ExpressionStatement":
		return codeGenerator(*n.expression) + ";"

	case "CallExpression":
		var ra []string
		c := codeGenerator(*n.callee)

		for _, no := range *n.arguments {
			ra = append(ra, codeGenerator(no))
		}

		r := strings.Join(ra, ", ")
		return c + "(" + r + ")"

	case "Identifier":
		return n.name

	case "NumberLiteral":
		return n.value

	default:
		log.Fatal("err")
		return ""
	}
}

func main() {
	tokens := tokenizer("(add 2 (subtract 10 5))")
	ast    := parser(tokens)
	nast   := transformer(ast)
	out    := codeGenerator(node(nast))

	fmt.Println(out)
}
