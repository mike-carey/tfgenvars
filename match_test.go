package main_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mcarey-solstice/tfgenvars"
)

var _ = Describe("tfgenvars", func() {
	Describe("Empty blocks", func() {
		It("Should pull an empty variable block", func() {
			var empty_block = string(`variable "empty_block" {}`)

			matches := CollectVariablesFromText(empty_block)

			Expect(matches).To(Equal([]string{empty_block}))
		})

		It("Should pull multiple empty variable blocks", func() {
			var empty_blocks = []string{`variable "empty_block_1" {}`, `variable "empty_block_2" {}`, `variable "empty_block_3" {}`}

			matches := CollectVariablesFromText(fmt.Sprintf("%s\n%s\n%s", empty_blocks[0], empty_blocks[1], empty_blocks[2]))

			Expect(matches).To(Equal(empty_blocks))
		})

		It("Should pull multiple empty variable blocks", func() {
			var empty_blocks = []string{`variable "empty_block_1" {}`, `variable "empty_block_2" {}`, `variable "empty_block_3" {}`}
			var resource_block = `resource "resource_block_1" {}`

			matches := CollectVariablesFromText(fmt.Sprintf("%s\n%s\n%s\n%s", empty_blocks[0], empty_blocks[1], empty_blocks[2], resource_block))

			Expect(matches).To(Equal(empty_blocks))
		})
	})

	Describe("Map blocks", func() {
		It("Should pull maps properly", func() {
			var map_block = `variable "map_block" {\n\n  default = {\n  foo = "bar"\n  }\n}`

			matches := CollectVariablesFromText(map_block)

			Expect(matches).To(Equal([]string{map_block}))
		})

		It("Should pull nested maps properly", func() {
			var map_block = `variable "map_block" {\n\n  default = {\n  foo = {\n    bar = "baz"\n  }\n  }\n}`

			matches := CollectVariablesFromText(map_block)

			Expect(matches).To(Equal([]string{map_block}))
		})
	})

	Describe("Remove non-variables", func() {
		It("Should remove non-variables", func() {
			var resource_block = `resource "server" "name" {\n  option = true\n}`
			var variable_block = `variable "variable_name" {\n  default = "foo"\n}`

			Context("No variables", func() {
				matches := CollectVariablesFromText(fmt.Sprintf("%s\n", resource_block))
				Expect(matches).To(Equal([]string{}))
			})

			Context("Variable after", func() {
				matches := CollectVariablesFromText(fmt.Sprintf("%s\n%s", resource_block, variable_block))
				Expect(matches).To(Equal([]string{variable_block}))
			})

			Context("Variable before", func() {
				matches := CollectVariablesFromText(fmt.Sprintf("%s\n%s", variable_block, resource_block))
				Expect(matches).To(Equal([]string{variable_block}))
			})

			Context("Variable before and after", func() {
				matches := CollectVariablesFromText(fmt.Sprintf("%s\n%s\n%s", variable_block, resource_block, variable_block))
				Expect(matches).To(Equal([]string{variable_block, variable_block}))
			})
		})
	})
})
