package tfgenvars

import (
	"bufio"
	"os"
	"errors"
	"strings"
)

var (
	BackwardsMovementError = errors.New("Moving backwards in buffer is prohibited")
	IndexOutOfBoundsError = errors.New("Index out of bounds")
)

type Buffer struct {
	Reader *bufio.Reader
	Position int
}

func NewBufferFromText(text string) *Buffer {
	b := Buffer{}
	b.Reader = bufio.NewReader(strings.NewReader(text))
	b.Position = 0
	return &b
}

func NewBuffer(file *os.File) *Buffer {
	b := Buffer{}
	b.Reader = bufio.NewReader(file)
	b.Position = 0
	return &b
}

// func (b *Buffer) Find(search string) (string, error) {
// 	str, err := b.Reader.ReadString(search)
// 	s := string(str)
//
// 	b.Position += len(s)
// 	return s, err
// }

func (b *Buffer) MoveTo(location int) (string, error) {
	diff := location - b.Position
	if diff == 0 {
		return "", nil
	} else if diff < 0 {
		return "", BackwardsMovementError
	} else if b.Reader.Size() < b.Position + diff {
		return "", IndexOutOfBoundsError
	}

	str := ""
	for i := 0; i < diff; i++ {
		s, _ := b.Next()
		str += s
	}
	return str, nil
}

func (b *Buffer) Next() (string, error) {
	c, _, err := b.Reader.ReadRune()
	b.Position += 1

	return string(c), err
}
