package main

import (
	"io"
	"io/ioutil"
	"regexp"
)

const VARIABLE_DECLARATION = "variable\\s+\"\\w+\"\\s+{"

func CollectVariablesFromFile(filename string) []string {
	s, e := ioutil.ReadFile(filename)
	if e != nil {
		panic(e)
	}

	return CollectVariablesFromText(string(s))
}

func CollectVariablesFromText(text string) []string {
	buffer := NewBufferFromText(text)

	regex := regexp.MustCompile(VARIABLE_DECLARATION)
	indices := regex.FindAllStringIndex(text, -1)

	pos := []int{}
	for _, i := range indices {
		pos = append(pos, i[0])
	}

	vars := CollectVariables(buffer, pos)

	return vars
}

func CollectVariables(buffer *Buffer, positions []int) []string {
	vars := []string{}
	for _, i := range positions {
		_, e := buffer.MoveTo(i)
		if e != nil {
			panic(e)
		}

		vars = append(vars, CollectVariableBlock(buffer))
	}

	return vars
}

func CollectVariableBlock(buffer *Buffer) string {
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
