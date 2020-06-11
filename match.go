package tfgenvars

import (
	"io"
	"io/ioutil"
	"regexp"
)

const (
	OutputDeclaration = "(?m)^output\\s+\"(\\w+)\"\\s+{"
	VariableDeclaration = "(?m)^variable\\s+\"(\\w+)\"\\s+{"
)

type Collector interface {
	CollectFromFile(filename string) map[string]string
	CollectFromText(text string) map[string]string
}

type CollectorImpl struct {
	regex *regexp.Regexp
}

func NewCollector(regex string) Collector {
	return &CollectorImpl{
		regex: regexp.MustCompile(regex),
	}
}

func (c *CollectorImpl) CollectFromFile(filename string) map[string]string {
	s, e := ioutil.ReadFile(filename)
	if e != nil {
		panic(e)
	}

	return c.CollectFromText(string(s))
}

func (c *CollectorImpl) CollectFromText(text string) map[string]string {
	buffer := NewBufferFromText(text)

	indices := c.regex.FindAllSubmatchIndex([]byte(text), -1)

	pos := make(map[string]int, 0)
	for _, i := range indices {
		// Indices: {
		//   0: first position found by regex,
		//   1: end position found by regex,
		//   2: first position of submatch,
		//   3: end position of submatch
		// }
		pos[text[i[2]:i[3]]] = i[1] - 1
	}

	// ps := make(positions, len(pos))
	// for key, index := range pos {
	// 	ps = append(ps, position{Key: key, Position: index,})
	// }
	// ps.Sort()
	// pos = ps.Map()

	vars := c.collect(buffer, pos)

	return vars
}

func (c *CollectorImpl) collect(buffer *Buffer, positions map[string]int) map[string]string {
	vars := make(map[string]string, len(positions))
	for key, i := range positions {
		logf("Moving to position: %d", i)
		_, e := buffer.MoveTo(i)
		if e != nil {
			panic(e)
		}

		vars[key] = c.collectBlock(buffer)
	}

	return vars
}

func (c *CollectorImpl) collectBlock(buffer *Buffer) string {
	str := ""
	v := 0
	b := false
	for {
		c, e := buffer.Next()
		if e == io.EOF {
			panic(e)
		}

		str += c

		switch c {
		case "{": {
			v += 1
		}
		case "}": {
			v -= 1
			if v == 0 {
				b = true
			}
		}
		}

		if b == true {
			break
		}
	}

	return str
}
