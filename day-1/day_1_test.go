package day1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	day1 "github.com/prycey77/2019/day-1"
)

var _ = Describe("Day 1", func() {
	It("Returns 2 when given 12", func() {
		Expect(day1.Counter_upper(12)).To(Equal(2))
	})
})
