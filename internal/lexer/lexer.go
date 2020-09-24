package lexer

import (
	"strconv"
)

type Lexer struct {
	Data  string
	index int
}

func New(input string) *Lexer {
	return &Lexer{
		Data:  input,
		index: 0,
	}
}

func (l *Lexer) Current() byte {
	return l.Data[l.index]
}

func (l *Lexer) Next() byte {
	l.index++

	return l.Current()
}

func (l *Lexer) Peek() byte {
	if len(l.Data) <= l.index+1 {
		return l.Data[len(l.Data)-1]
	}

	return l.Data[l.index+1]
}

func (l *Lexer) EOF() bool {
	return len(l.Data) <= l.index
}

func (l *Lexer) NextEOF() bool {
	return len(l.Data) <= l.index+1
}

func (l *Lexer) Gobble() {
	for {
		if !IsWhitespace(l.Current()) {
			return
		}
		l.index++
	}
}

func (l *Lexer) Print() string {
	return l.Data[l.index:]
}

func (l *Lexer) Advance(i int) byte {
	l.index += i
	return l.Current()
}

func (l *Lexer) Backup() byte {
	l.index--
	return l.Current()
}

func (l *Lexer) Snapshot() string {
	start := l.index - 15
	if start <= 0 {
		start = 0
	}

	end := l.index + 15
	if end >= len(l.Data) {
		end = len(l.Data)
	}

	return string(l.Data[start:l.index]) + "!>" + string(l.Data[l.index]) + "<!" + string(l.Data[l.index+1:end])
}

func (l *Lexer) Word() string {
	var c byte

	l.Gobble()
	c = l.Current()

	var word string
	for {

		if c == ' ' || c == '\n' || c == '\r' || l.EOF() {
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
	for IsWhitespace(l.Current()) {
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
	for IsWhitespace(l.Current()) {
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

	if !IsDigit(c) {
		panic("Fread_number: bad format: " + string(l.Snapshot()))
	}

	var number int
	for IsDigit(c) {
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

func IsWhitespace(c byte) bool {
	return c == ' ' || c == '\n' || c == '\t' || c == '\r'
}

func IsDigit(c byte) bool {
	_, err := strconv.Atoi(string(c))
	return err == nil
}
