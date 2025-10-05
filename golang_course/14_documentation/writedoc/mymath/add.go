// Package writedoc provides ACME inc math solutions.
package writedoc

// Sum of number
func Sum(no ...int) int {

	sum := 0
	for i := range no {
		sum = i + sum
	}
	return sum
}
