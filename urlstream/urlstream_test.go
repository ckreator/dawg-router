package urlstream

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStreamifyBasicRoute(t *testing.T) {
	str := "/abc/de/f"
	input := makeIterator(str)
	i := 0
	for !input.EOF() {
		fmt.Printf("GOT CHAR: %s | %s\n", input.Peek(), string(str[i]))
		assert.Equal(t, input.Next(), string(str[i]), "Should be the same string as at index")
		i++
	}
}

func TestTokenizeParamRoute(t *testing.T) {
	assert := assert.New(t)
	str := "/a/:a_id"
	stream := makeIterator(str)
	input := tokenizeStream(*stream)

	tok := input.Next()
	fmt.Println("Testing /", tok.Value)
	assert.Equal("/", tok.Value, "Value should equal slash")
	tok = input.Next()

	fmt.Println("Testing a", tok)
	assert.Equal("a", tok.Value, "Value should equal 'a'")
	tok = input.Next()

	fmt.Println("Testing second /", tok)
	assert.Equal("/", tok.Value, "Value should equal slash")
	tok = input.Next()

	fmt.Println("Testing a_id", tok.Value)
	assert.Equal("a_id", tok.Value, "Value should equal a_id")
}
