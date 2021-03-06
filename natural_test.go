package strcmp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test__string_equals_itself(t *testing.T) {
	assertEqual(t, "", "")
	assertEqual(t, "x", "x")
	assertEqual(t, "1", "1")
	assertEqual(t, "x1", "x1")
	assertEqual(t, "1x", "1x")
	assertEqual(t, "xx", "xx")
	assertEqual(t, "11", "11")
	assertEqual(t, "x11", "x11")
	assertEqual(t, "11x", "11x")
	assertEqual(t, "x11x", "x11x")
}

func Test__empty_is_less_than_any_nonempty(t *testing.T) {
	assertLessThan(t, "", "x")
	assertLessThan(t, "", "1")
	assertLessThan(t, "", "x1")
	assertLessThan(t, "", "1x")
}

func Test__single_letter_comparison(t *testing.T) {
	assertLessThan(t, "x", "y")
}

func Test__single_digit_comparison(t *testing.T) {
	assertLessThan(t, "0", "1")
}

func Test__number_comparison(t *testing.T) {
	assertLessThan(t, "0", "1")
	assertLessThan(t, "00", "1")
	assertLessThan(t, "0", "10")
	assertLessThan(t, "00", "10")

	assertLessThan(t, "2", "10")
	assertLessThan(t, "10", "20")
}

func Test__number_comparison_with_leading_zero(t *testing.T) {
	assertLessThan(t, "0", "00")
	assertLessThan(t, "1", "01")
	assertLessThan(t, "1", "001")
	assertLessThan(t, "01", "001")
}

func Test__number_comparison_with_trailing_zero(t *testing.T) {
	assertLessThan(t, "0", "00")
	assertLessThan(t, "1", "10")
	assertLessThan(t, "1", "100")
	assertLessThan(t, "10", "100")
}

func Test__number_comparison_after_letter_prefix(t *testing.T) {
	// Digit is compared to letter (note that numbers are always "larger").
	assertLessThan(t, "x2", "10")
	assertLessThan(t, "x10", "2")

	// Numbers are compared.
	assertLessThan(t, "x2", "x10")
	assertLessThan(t, "x12", "x21")
}

func Test__number_comparison_before_letter_suffix(t *testing.T) {
	assertLessThan(t, "2x", "10")
	assertLessThan(t, "1", "10x")
	assertLessThan(t, "2", "10x")
	assertLessThan(t, "02", "10x")
	assertLessThan(t, "2x", "10x")
}

func Test__number_comparison_between_letters(t *testing.T) {
	assertLessThan(t, "x2x", "x10")
	assertLessThan(t, "x2", "x10x")
	assertLessThan(t, "x2x", "x10x")
	assertLessThan(t, "x1x", "x01x")
}

func Test__number_comparison_after_equal_number(t *testing.T) {
	assertLessThan(t, "x1x", "x10")
	assertLessThan(t, "x1X", "x10")
	assertLessThan(t, "x1 ", "x10")
}

func Test__number_with_different_suffix(t *testing.T) {
	assertLessThan(t, "x2x", "x2y")
}

func Test__number_comparison_with_same_suffix(t *testing.T) {
	assertLessThan(t, "0-", "1-")
	assertLessThan(t, "0x", "1x")
	assertLessThan(t, "0xy", "1xy")
	assertLessThan(t, "1x", "010x")
	assertLessThan(t, "0x0", "1x0")
}

func Test__number_comparison_with_different_suffix(t *testing.T) {
	assertLessThan(t, "0xa", "1xb")
	assertLessThan(t, "0xb", "1xa")
	assertLessThan(t, "1xa", "010xb")

	assertLessThan(t, "0x1", "1x0")
	assertLessThan(t, "0a1", "1b0")
	assertLessThan(t, "0b1", "1a0")
}

func Test__letter_comparison(t *testing.T) {
	assertLessThan(t, "a1", "b0")
	assertLessThan(t, "ab", "ba")
}

func assertEqual(t *testing.T, left, right string) {
	r := Natural(left, right)
	assert.Equal(t, 0, r, "Natural(%s, %s)", left, right)
}

func assertLessThan(t *testing.T, left, right string) {
	r1 := Natural(left, right)
	assert.Equal(t, -1, r1, "Natural(%s, %s)", left, right)

	r2 := Natural(right, left)
	assert.Equal(t, 1, r2, "Natural(%s, %s)", right, left)
}
