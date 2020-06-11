package tfgenvars_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/tfgenvars"
)

var _ = Describe("Pointers", func() {

	It("Should sort pointers", func() {
		foo := Pointer{
			Name: "foo",
			Position: 101,
		}
		bar := Pointer{
			Name: "bar",
			Position: 100,
		}
		pointers := Pointers{foo, bar}
		pointers.Sort()

		Expect(pointers).To(Equal(Pointers{bar, foo}))
	})

})
