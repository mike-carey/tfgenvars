package tfgenvars_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/tfgenvars"
)

const (
	variableType  = "variable"
	outputType = "output"
)

type block struct {
	Type string
	Name string
	Code string
}

func (b *block) String() string {
	return fmt.Sprintf("%s \"%s\" %s", b.Type, b.Name, b.Code)
}

func (b *block) Map() map[string]string {
	return map[string]string { b.Name: b.Code }
}

func newBlock(theType string, name string, code string) *block {
	return &block{
		Type: theType,
		Name: name,
		Code: code,
	}
}

func newVariableBlock(name string, code string) *block {
	return newBlock(variableType, name, code)
}

func newOutputBlock(name string, code string) *block {
	return newBlock(outputType, name, code)
}

type blocks []*block

func (bs blocks) String() string {
	s := ""
	for _, b := range bs {
		s += b.String() + "\n"
	}
	return s
}

func (bs blocks) Map() map[string]string {
	m := make(map[string]string, len(bs))
	for _, b := range bs {
		m[b.Name] = b.Code
	}
	return m
}

var _ = Describe("tfgenvars", func() {
	Describe("Empty blocks", func() {
		var (
			collector Collector
		)

		BeforeEach(func() {
			collector = NewCollector(VariableDeclaration)
		})

		It("Should pull an empty variable block", func() {
			emptyBlock := newVariableBlock("empty_block", "{}")
			matches := collector.CollectFromText(emptyBlock.String())
			Expect(matches).To(Equal(emptyBlock.Map()))
		})

		It("Should pull multiple empty variable blocks", func() {
			emptyBlocks := blocks{
				newVariableBlock("empty_block_0", "{}"),
				newVariableBlock("empty_block_1", "{}"),
				newVariableBlock("empty_block_2", "{}"),
			}

			matches := collector.CollectFromText(emptyBlocks.String())

			Expect(matches).To(Equal(emptyBlocks.Map()))
		})

		It("Should pull multiple empty variable blocks", func() {
			emptyBlocks := []string {`variable "empty_block_1" {}`, `variable "empty_block_2" {}`, `variable "empty_block_3" {}` }
			resourceBlock := `resource "resource_block_1" {}`

			textBlock := fmt.Sprintf("%s\n%s\n%s\n%s", emptyBlocks[0], emptyBlocks[1], emptyBlocks[2], resourceBlock)
			expectedRes := make(map[string]string, len(emptyBlocks))
			for i, _ := range emptyBlocks {
				expectedRes[fmt.Sprintf("empty_block_%d", i + 1)] = "{}"
			}

			matches := collector.CollectFromText(textBlock)

			Expect(matches).To(Equal(expectedRes))
		})
	})

	// Describe("Map blocks", func() {
	// 	It("Should pull maps properly", func() {
	// 		var map_block = `variable "map_block" {\n\n  default = {\n  foo = "bar"\n  }\n}`
	//
	// 		matches := CollectVariablesFromText(map_block)
	//
	// 		Expect(matches).To(Equal([]string{map_block}))
	// 	})
	//
	// 	It("Should pull nested maps properly", func() {
	// 		var map_block = `variable "map_block" {\n\n  default = {\n  foo = {\n    bar = "baz"\n  }\n  }\n}`
	//
	// 		matches := CollectVariablesFromText(map_block)
	//
	// 		Expect(matches).To(Equal([]string{map_block}))
	// 	})
	// })

// 	Describe("Remove non-variables", func() {
// 		It("Should remove non-variables", func() {
// 			var resource_block = `resource "server" "name" {\n  option = true\n}`
// 			var variable_block = `variable "variable_name" {\n  default = "foo"\n}`
//
// 			Context("No variables", func() {
// 				matches := CollectVariablesFromText(fmt.Sprintf("%s\n", resource_block))
// 				Expect(matches).To(Equal([]string{}))
// 			})
//
// 			Context("Variable after", func() {
// 				matches := CollectVariablesFromText(fmt.Sprintf("%s\n%s", resource_block, variable_block))
// 				Expect(matches).To(Equal([]string{variable_block}))
// 			})
//
// 			Context("Variable before", func() {
// 				matches := CollectVariablesFromText(fmt.Sprintf("%s\n%s", variable_block, resource_block))
// 				Expect(matches).To(Equal([]string{variable_block}))
// 			})
//
// 			Context("Variable before and after", func() {
// 				matches := CollectVariablesFromText(fmt.Sprintf("%s\n%s\n%s", variable_block, resource_block, variable_block))
// 				Expect(matches).To(Equal([]string{variable_block, variable_block}))
// 			})
// 		})
// 	})
//
// 	Describe("Edge cases", func() {
// 		It("Should not pick up a resource named variable", func() {
// 			var resource_block = `
// resource "variable" "foo" {
//   value = "bar"
// }`
// 			matches := CollectVariablesFromText(resource_block)
// 			Expect(matches).To(Equal([]string{}))
// 		 })
//
// 		 It("Should not pick up a commented out variable", func() {
// 			 var comment_block = `
// // variable "foo" {}
// }`
// 		      matches := CollectVariablesFromText(comment_block)
// 		      Expect(matches).To(Equal([]string{}))
// 		 })
//
// 	})
})
