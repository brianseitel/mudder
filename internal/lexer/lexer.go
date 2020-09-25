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

func New(input string) *Lexer {
	return &Lexer{
		data:  input,
		index: 0,
	}
}

func (l *Lexer) setIndex(i int) *Lexer {
	l.index = i
	return l
}

func (l *Lexer) Current() byte {
	return l.data[l.index]
}

func (l *Lexer) Next() byte {
	l.index++

	return l.Current()
}

func (l *Lexer) Peek() byte {
	if len(l.data) <= l.index+1 {
		return l.data[len(l.data)-1]
	}

	return l.data[l.index+1]
}

func (l *Lexer) EOF() bool {
	return len(l.data) <= l.index+1
}

func (l *Lexer) NextEOF() bool {
	return len(l.data) <= l.index+1
}

func (l *Lexer) Gobble() {
	for {
		if !isWhitespace(l.Current()) {
			return
		}
		l.index++
	}
}

func (l *Lexer) Print() string {
	return l.data[l.index:]
}

func (l *Lexer) Advance(i int) byte {
	l.index += i
	if l.index >= len(l.data) {
		l.index = len(l.data) - 1
	}
	return l.Current()
}

func (l *Lexer) Backup() byte {
	l.index--
	return l.Current()
}

func (l *Lexer) Jump(target string) error {
	start := strings.Index(l.data, target)
	if start == -1 {
		return errors.New("target not found")
	}

	// found it! let's skip over...
	l.Advance(start + len(target))
	return nil
}

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
		if c == ' ' || c == '\n' || c == '\r' {
			return word
		}

		word += string(c)
		if l.NextEOF() {
			return word
		}
		c = l.Next()
	}
}

func (l *Lexer) Letter() string {
	l.Gobble()

	out := l.Current() // get the Current letter
	l.Advance(1)       // Gobble it up and move on to the Next one
	return string(out)
}

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

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\n' || c == '\t' || c == '\r'
}

func isDigit(c byte) bool {
	_, err := strconv.Atoi(string(c))
	return err == nil
}
