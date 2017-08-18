package jumpSearch

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Just Search Teste Cases")
}

var _ = Describe("jump Search", func() {
	It("search the elements", func() {
		Expect(Search([]int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 12)).To(Equal(1))
		Expect(Search([]int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 112)).To(Equal(-1))

		Expect(Search([]int{29897, 34834, 43379, 46179, 60064, 64337, 71814, 73286, 83304, 96122}, 43379)).To(Equal(2))
		Expect(Search([]int{29897, 34834, 43379, 46179, 60064, 64337, 71814, 73286, 83304, 96122}, 118)).To(Equal(-1))

		Expect(Search([]int{11, 12}, 11)).To(Equal(0))
		Expect(Search([]int{11, 12}, 10)).To(Equal(-1))

		Expect(Search([]int{11}, 11)).To(Equal(0))
		Expect(Search([]int{11}, 10)).To(Equal(-1))

		Expect(Search([]int{20, 19, 18, 17, 16, 15, 14, 13, 12, 10}, 12)).To(Equal(-1))
		Expect(Search([]int{13, 12, 10}, 12)).To(Equal(-1))

		Expect(Search([]int{43379, 46179, 64337, 83304, 73286, 29897, 60064, 96122, 34834, 71814}, 60064)).To(Equal(-1))
		Expect(Search([]int{43379, 46179, 64337}, 60064)).To(Equal(-1))
	})
})
