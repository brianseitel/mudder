package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexerCurrent(t *testing.T) {
	l := New("this is a test")

	assert.Equal(t, byte('t'), l.Current())
}

func TestLexerNext(t *testing.T) {
	l := New("this is a test")

	assert.Equal(t, byte('h'), l.Next())
	assert.Equal(t, 1, l.index)
}

func TestLexerPrint(t *testing.T) {
	l := New("this is a test")

	assert.Equal(t, "this is a test", l.Print())
}

func TestLexerPeek(t *testing.T) {
	l := New("this is a test")

	assert.Equal(t, byte('h'), l.Peek())
	assert.Equal(t, 0, l.index)

	l.index = 100
	assert.Equal(t, byte('t'), l.Peek())
}

func TestLexerJump(t *testing.T) {
	l := New("#BAR baz")

	err := l.Jump("baz")

	assert.Nil(t, err)
	assert.Equal(t, 7, l.index)
	assert.Error(t, l.Jump("should be an error"))
}

func TestLexerEOF(t *testing.T) {
	l := New("foo")
	l.index = 3 // go to the end

	assert.True(t, l.EOF())
}

func TestGobble(t *testing.T) {
	l := New("        this is a big fat foo")

	assert.Equal(t, 0, l.index)
	l.Gobble()
	assert.Equal(t, 8, l.index)
}

func TestAdvance(t *testing.T) {
	l := New("this is a test")

	assert.Equal(t, byte('s'), l.Advance(3))
	assert.Equal(t, 3, l.index)
}

func TestBackup(t *testing.T) {
	l := New("this is a test")

	l.Advance(3)
	assert.Equal(t, byte('s'), l.Current())
	l.Backup()
	assert.Equal(t, byte('i'), l.Current())
	assert.Equal(t, 2, l.index)
}

func TestSnapshot(t *testing.T) {
	cases := []struct {
		Input    *Lexer
		Expected string
	}{
		{
			Input:    New(`it was many and many a year ago`).setIndex(16),
			Expected: "t was many and !>m<!any a year ago",
		},
		{
			Input:    New(`short`).setIndex(3),
			Expected: "sho!>r<!t",
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.Expected, c.Input.Snapshot())
	}
}

func TestWord(t *testing.T) {
	l := New("foo bar")

	assert.Equal(t, "foo", l.Word())
	assert.Equal(t, "bar", l.Word())
}

func TestLetter(t *testing.T) {
	l := New("S foo E bar")

	assert.Equal(t, "S", l.Letter())
	assert.Equal(t, "foo", l.Word())
	assert.Equal(t, "E", l.Letter())
	assert.Equal(t, "b", l.Letter())

	l = New("1d1+29999")

	assert.Equal(t, "1", l.Letter())
	assert.Equal(t, "d", l.Letter())
	assert.Equal(t, 1, l.Number())
	assert.Equal(t, 29999, l.Number())

}

func TestEOL(t *testing.T) {
	cases := []struct {
		Input    *Lexer
		Expected string
	}{
		{
			Input: New(`but soft what light through yonder window breaks
it is the east and juliet is the sun
arise fair sun and kill the envious moon
for it is sick and pale with grief
that thou her maid art far more fair than she`),
			Expected: "but soft what light through yonder window breaks",
		},
		{
			Input:    New(""),
			Expected: "",
		},
		{
			Input:    New("\n"),
			Expected: "",
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.Expected, c.Input.EOL())
	}

}

func TestString(t *testing.T) {
	cases := []struct {
		Input    *Lexer
		Expected string
	}{
		{
			Input:    New("there can be only one~"),
			Expected: "there can be only one",
		},
		{
			Input:    New("    buncha whitespace up front~"),
			Expected: "buncha whitespace up front",
		},
		{
			Input:    New("~"),
			Expected: "",
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.Expected, c.Input.String())
	}
}

func TestNumber(t *testing.T) {
	cases := []struct {
		Input    *Lexer
		Expected []int
	}{
		{
			Input:    New("25 15 1"),
			Expected: []int{25},
		},
		{
			Input:    New("    19 6"),
			Expected: []int{19},
		},
		{
			Input:    New("4|8"),
			Expected: []int{12},
		},
		{
			Input:    New("8|128"),
			Expected: []int{136},
		},
		{
			Input:    New("-16"),
			Expected: []int{-16},
		},
		{
			Input:    New("1 2"),
			Expected: []int{1, 2},
		},
		{
			Input:    New("+12345"),
			Expected: []int{12345},
		},
	}

	for _, c := range cases {
		for _, ex := range c.Expected {
			assert.Equal(t, ex, c.Input.Number())
		}
	}
}
