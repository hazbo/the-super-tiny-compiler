package main

import (
	"reflect"
	"testing"
)

const input = "(add 10 (subtract 10 6))"

func TestTokenizer(t *testing.T) {
	expected := []token{
		token{
			kind:  "paren",
			value: "(",
		},
		token{
			kind:  "name",
			value: "add",
		},
		token{
			kind:  "number",
			value: "10",
		},
		token{
			kind:  "paren",
			value: "(",
		},
		token{
			kind:  "name",
			value: "subtract",
		},
		token{
			kind:  "number",
			value: "10",
		},
		token{
			kind:  "number",
			value: "6",
		},
		token{
			kind:  "paren",
			value: ")",
		},
		token{
			kind:  "paren",
			value: ")",
		},
	}
	result := tokenizer(input)
	if !reflect.DeepEqual(result, expected) {
		t.Error("\nExpected:", expected, "\nGot:", result)
	}
}
