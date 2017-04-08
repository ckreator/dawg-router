package urlstream

import (
	"fmt"
	"log"
	"strings"
)

// Token is a generic token format
type Token struct {
	TokType string
	Value   string
}

// InputStream is a generic type for input streams
type InputStream struct {
	Next      func() string
	Peek      func() string
	EOF       func() bool
	Suffocate func(string, ...string)
}

// Tokenizer is a type that acts like an InputStream, but for tokens
type Tokenizer struct {
	Next      func() Token
	Peek      func() Token
	EOF       func() bool
	Suffocate func(string, ...string)
}

// TODO: make Iterator could be 'channelized' to increase processing speed of string
func makeIterator(input string) *InputStream {
	row := 1
	col := 1
	pos := 0
	ret := new(InputStream)
	splitted := strings.Split(input, "")
	var c string

	// let's roll with utf8 support
	/*splitted := make([]rune, 0)
	for i, w := 0, 0; i < len(input); i += w {
		runeValue, width := utf8.DecodeRuneInString(input[i:])
		//fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		splitted = append(splitted, runeValue)
		w = width
	}*/
	//fmt.Printf("%#U\n", ary[0])

	next := func() string {
		c = splitted[pos]
		//c = fmt.Sprintf("%c", splitted[pos])
		pos++
		if c == "\n" {
			col = 1
			row++
		} else {
			col++
		}
		return c
	}

	eof := func() bool {
		return pos >= len(splitted)
	}

	peek := func() string {
		if eof() {
			return ""
		}
		//c = fmt.Sprintf("%c", splitted[pos])
		return splitted[pos]
	}

	suffocate := func(msg string, list ...string) {
		beg := 0
		// maybe choose another number
		if pos > 20 {
			beg = pos - 20
		}

		log.Fatal(input[beg:pos], "\n                     ^\n", msg, fmt.Sprintf(" at row %d, col %d\n", row, col))
		panic(msg)
	}

	ret.Next = next
	ret.Peek = peek
	ret.EOF = eof
	ret.Suffocate = suffocate
	return ret
}

// tokenize the stream, actually just filter out params
func tokenizeStream(input InputStream) Tokenizer {
	pos := 0
	var currTok Token = Token{TokType: "Param", Value: ""}
	var tokens []Token = []Token{}
	mode := "SIMPLE"
	for !input.EOF() {
		nxt := input.Next()
		if nxt == ":" {
			mode = "PARAM"
			continue
		}
		// now do separate stuff for each mode
		if mode == "PARAM" {
			if nxt == "/" {
				mode = "SIMPLE"
				tokens = append(tokens, currTok)
				// clear value
				currTok.Value = ""
				tokens = append(tokens, Token{TokType: "Basic", Value: nxt})
			} else {
				currTok.Value += nxt
			}
		} else if mode == "SIMPLE" {
			tokens = append(tokens, Token{TokType: "Basic", Value: nxt})
		}
	}
	if currTok.Value != "" {
		tokens = append(tokens, currTok)
	}
	// create tokenizer object
	ret := Tokenizer{}

	eof := func() bool {
		return pos >= len(tokens)
	}

	next := func() Token {
		if eof() {
			return Token{}
		}
		pos++
		//fmt.Println("GETTING AT:", pos-1, len(tokens))
		return tokens[pos-1]
	}

	peek := func() Token {
		if eof() {
			return Token{}
		}
		return tokens[pos]
	}

	suffocate := func(msg string, list ...string) {
		fmt.Println("NOOOOOOOOO, Tokenizer dead, he dead!!!")
	}

	ret.Next = next
	ret.Peek = peek
	ret.EOF = eof
	ret.Suffocate = suffocate
	return ret
}
