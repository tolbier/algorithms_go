package multiplier

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Internal Tests", func() {
	Context("multiplier.splitHalves()", func() {
		It("should separate string returning the two halves and their length", func() {
			inputs := []string{
				"1234",
				"123",
			}
			results := [][]interface{}{
				{"12", "34"},
				{"01", "23"},
			}
			for idx, result := range results {
				a, b := splitHalves(inputs[idx])
				Expect(a).Should(Equal(result[0]))
				Expect(b).Should(Equal(result[1]))
			}
		})
	})
	Context("multiplier.normalizeZeroes()", func() {
		It("should return strings with same length, left-zeroing", func() {
			inputs := [][]string{
				{"12", "34"},
				{"1", "23"},
				{"12", "3"},
				{"12123", "3"},
				{"3", "12343"},
			}
			results := [][]interface{}{
				{"12", "34"},
				{"01", "23"},
				{"12", "03"},
				{"12123", "00003"},
				{"00003", "12343"},
			}
			for idx, result := range results {
				a, b := normalizeZeroes(inputs[idx][0], inputs[idx][1])
				Expect(a).Should(Equal(result[0]))
				Expect(b).Should(Equal(result[1]))
			}
		})
	})
	Context("multiplier.digitMultiply()", func() {
		It("should multiply digits", func() {
			inputs := [][]string{
				{"1", "4"},
				{"2", "3"},
				{"9", "9"},
			}
			results := []string{"4", "6", "81"}
			for idx, result := range results {
				a := digitMultiply(inputs[idx][0], inputs[idx][1])
				Expect(a).Should(Equal(result))
			}
		})
	})
	Context("multiplier.sumString()", func() {
		It("should sum string numbers", func() {
			inputs := [][]string{
				{"1", "4"},
				{"2123", "3234556"},
				{"934534", "9999"},
				{"9999", "9999"},
				{"03", "08"},
			}
			results := []string{"5", "3236679", "944533", "19998", "11"}
			for idx, result := range results {
				a := sumString(inputs[idx][0], inputs[idx][1])
				Expect(a).Should(Equal(result))
			}
		})
	})
	Context("multiplier.substractString()", func() {
		It("should substract string numbers", func() {
			inputs := [][]string{
				{"4", "1"},
				{"21230394", "3234556"},
				{"934534", "9999"},
				{"11111", "9999"},
			}
			results := []string{"3", "17995838", "924535", "1112"}
			for idx, result := range results {
				a := substractString(inputs[idx][0], inputs[idx][1])
				Expect(a).Should(Equal(result))
			}
		})
	})

	Context("multiplier.bypow10()", func() {
		It("should multiply by Pow10(n)", func() {
			inputs := [][]interface{}{
				{"4", 0},
				{"23", 1},
				{"43", 2},
				{"44", 21},
			}
			results := []string{"4", "230", "4300", "44000000000000000000000"}
			for idx, result := range results {
				a := bypow10(inputs[idx][0].(string), inputs[idx][1].(int))
				Expect(a).Should(Equal(result))
			}
		})
	})
	Context("multiplier.removeZeroes()", func() {
		It("should remove all left Zeroes", func() {
			inputs := []string{"4", "00123", "00123123112312312312323123123123123123123"}

			results := []string{"4", "123", "123123112312312312323123123123123123123"}
			for idx, result := range results {
				a := removeZeroes(inputs[idx])
				Expect(a).Should(Equal(result))
			}
		})
	})
})
