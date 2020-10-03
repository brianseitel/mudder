package lexer

import (
	"errors"
	"strconv"
	"strings"
)

type Lexer struct {
	data  string
	index int
}

// New makes a brand new Lexer. Defaults to index 0.
func New(input string) *Lexer {
	return &Lexer{
		data:  input,
		index: 0,
	}
}

// setIndex is used only for testing to set an index
// to an arbitrary point
func (l *Lexer) setIndex(i int) *Lexer {
	l.index = i
	return l
}

// Current returns the current index. This does
// not check to see if it's a valid index. Instead
// it just panics. Whoops.
func (l *Lexer) Current() byte {
	return l.data[l.index]
}

// Next returns the next byte in the sequence. It
// does not check whether the next byte is valid.
func (l *Lexer) Next() byte {
	l.index++

	return l.Current()
}

// Peek allows the caller to see what the next byte
// is without advancing the cursor. If it tries
// to peek past the end of the lexer body, it
// will return the last byte instead.
func (l *Lexer) Peek() byte {
	if len(l.data) <= l.index+1 {
		return l.data[len(l.data)-1]
	}

	return l.data[l.index+1]
}

// EOF checks whether the current byte is past
// the end of the file
func (l *Lexer) EOF() bool {
	return len(l.data) <= l.index
}

// NextEOF determines whether next byte is past
// the end of the file
func (l *Lexer) NextEOF() bool {
	return len(l.data) <= l.index+1
}

// Gobble eats any whitespace characters, advancing
// the index until it reaches a non-whitespace character.
func (l *Lexer) Gobble() {
	for {
		if !isWhitespace(l.Current()) {
			return
		}
		l.index++
	}
}

// Print returns the rest of the body from the index forward.
func (l *Lexer) Print() string {
	return l.data[l.index:]
}

// Advance moves the cursor up by an arbitrary amount. if
// the index goes past the length of the body, it will
// only advance to the last byte.
func (l *Lexer) Advance(i int) byte {
	l.index += i
	if l.index >= len(l.data) {
		l.index = len(l.data) - 1
	}
	return l.Current()
}

// Backup will decrement the index and
// return the current byte.
func (l *Lexer) Backup() byte {
	l.index--
	return l.Current()
}

// Jump will jump to a particular string point in
// the body, advancing the index to that point. If
// the substring is not found, it will return an
// error.
func (l *Lexer) Jump(target string) error {
	start := strings.Index(l.data, target)
	if start == -1 {
		return errors.New("target not found")
	}

	// found it! let's skip over...
	l.Advance(start + len(target))
	return nil
}

// Snapshot is used for debugging and will return
// up to 15 bytes on either side of the current byte
// index.
func (l *Lexer) Snapshot() string {
	start := l.index - 15
	if start <= 0 {
		start = 0
	}

	end := l.index + 15
	if end >= len(l.data) {
		end = len(l.data)
	}

	return string(l.data[start:l.index]) + "!>" + string(l.data[l.index]) + "<!" + string(l.data[l.index+1:end])
}

// Word finds the current word slice in the body.
// If it finds whitespace, it ends.
func (l *Lexer) Word() string {
	var c byte

	l.Gobble()
	c = l.Current()

	var word string
	for {
		if l.EOF() {
			word += string(c)
			return word
		}
		word += string(c)
		if l.NextEOF() {
			l.data += " "
			return word
		}
		c = l.Next()
		if c == ' ' || c == '\n' || c == '\r' {
			return word
		}

	}
}

// Letter returns the next non-whitespace character
func (l *Lexer) Letter() string {
	l.Gobble()

	out := l.Current() // get the Current letter
	l.Advance(1)       // Gobble it up and move on to the Next one
	return string(out)
}

// EOL returns each byte up until an end-of-line character.
// If it reaches the end of the file, it returns nothing.
func (l *Lexer) EOL() string {
	var c byte
	if l.EOF() {
		return ""
	}
	if l.Current() == '\n' || l.Current() == '\r' {
		return ""
	}
	for isWhitespace(l.Current()) {
		l.Next()
	}

	var output string
	for {
		c = l.Current()
		if c == '\n' || c == '\r' {
			return output
		}

		output += string(c)
		if l.NextEOF() {
			return output
		}
		c = l.Next()

	}
}

// String returns a tilde-terminated (~) string, ignoring
// any leading whitespace.
func (l *Lexer) String() string {
	var c byte
	for isWhitespace(l.Current()) {
		l.Next()
	}

	if l.Current() == '~' {
		if !l.NextEOF() {
			l.Next() // Gobble it up
		}
		return "" // no strang here!
	}

	var output string
	for {
		c = l.Current()
		if c == '~' {
			if !l.NextEOF() {
				l.Next() // pop off the tilde
			}
			return output
		}
		output += string(c)
		if l.NextEOF() {
			return output
		}
		c = l.Next()
	}
}

// Number returns thenext integer, positive or negative,
// in the body. It will read either a leading + or - sign.
// It supports combinations of integers separated by
// pipes (|). This is useful for flags. If it cannot find
// a number after clearing leading whitespace, it will panic.
func (l *Lexer) Number() int {
	sign := false
	var c byte
	l.Gobble()

	c = l.Current()

	if c == '+' {
		c = l.Next()
	} else if c == '-' {
		sign = true
		c = l.Next()
	}

	if !isDigit(c) {
		panic("Fread_number: bad format: " + string(l.Snapshot()))
	}

	var number int
	for isDigit(c) {
		num, _ := strconv.Atoi(string(c))
		number = number*10 + num
		if l.NextEOF() {
			break
		}
		c = l.Next()
	}

	if sign {
		number = 0 - number
	}

	if c == '|' {
		l.Next()
		number += l.Number()
	}
	return number
}

func SplitArgs(input string) (string, string) {
	words := strings.Fields(input)
	if len(words) == 1 {
		return input, ""
	} else if len(words) > 1 {
		return words[0], strings.Join(words[1:], " ")
	}
	return "", ""
}

// isWhitespace is used internally to define whitespace values.
func isWhitespace(c byte) bool {
	return c == ' ' || c == '\n' || c == '\t' || c == '\r'
}

// isDigit is used internally to determine whether a byte is
// a numeric digit or not.
func isDigit(c byte) bool {
	_, err := strconv.Atoi(string(c))
	return err == nil
}
