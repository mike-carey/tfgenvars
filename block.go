package tfgenvars

import (

)

const (
	VariableType = "variable"
	OutputType = "output"
)

type Block struct {
	Type string
	Name string
	Code string
}

func (b *Block) String() string {
	return fmt.Sprintf("%s \"%s\" %s", b.Type, b.Name, b.Code)
}

func (b *Block) Map() map[string]string {
	return map[string]string { b.Name: b.Code }
}

func NewBlock(theType string, name string, code string) *block {
	return &block{
		Type: theType,
		Name: name,
		Code: code,
	}
}

func NewVariableBlock(name string, code string) *block {
	return NewBlock(VariableType, name, code)
}

func NewOutputBlock(name string, code string) *block {
	return NewBlock(OutputType, name, code)
}

type Blocks []*Block

func (bs Blocks) String() string {
	s := ""
	for _, b := range bs {
		s += b.String() + "\n"
	}
	return s
}

func (bs Blocks) Map() map[string]string {
	m := make(map[string]string, len(bs))
	for _, b := range bs {
		m[b.Name] = b.Code
	}
	return m
}
