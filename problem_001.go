package main

/*
https://projecteuler.net/problem=1

Problem 1:
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

func P0001(cap int, multiplesOf []int) int {
	sum := 0
	for number := 1; number < cap; number++ {
		for _, multipleOf := range multiplesOf {
			if number%multipleOf == 0 {
				sum += number

				// break so that we do not match more than one multpiple against "number"
				break
			}
		}
	}

	return sum
}
