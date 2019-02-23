package tfgenvars_test

import (
	"os"
	"io/ioutil"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gbytes"

	. "github.com/mike-carey/tfgenvars"
)

func TestTfgenvars(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tfgenvars Suite")
}


var _ = Describe("Integration", func() {

	It("Should pass through stdin if no arguments were passed", func() {
		buffer := gbytes.BufferWithBytes([]byte(`
variable "foo" {
  default = "bar"
}

resource "baz" {
  default = "foo"
}
`))
		err := Run(buffer, buffer, []string{})
		Expect(err).To(BeNil())

		Expect(buffer).To(gbytes.Say(string(`variable "foo" {
  default = "bar"
}
`)))
	})

	It("Should collect files in arguments", func() {
		// Write file
		f, e := ioutil.TempFile(os.TempDir(), "")
		Expect(e).To(BeNil())

		f.Write([]byte(`
variable "foo" {}

resource "bar" "" {
	key1 = true
	key2 = {
		k = "v"
	}
}
`))
		buffer := gbytes.NewBuffer()

		e = Run(buffer, buffer, []string{f.Name()})
		Expect(e).To(BeNil())

		Expect(buffer).To(gbytes.Say(string(`variable "foo" {}
`)))
	})

})
