package workdir

const readme = `
Leetcode solutions
=======================================

# {{SolvedProblemsCount}} Solved Problems

|**No.**|**Problem**|**Solution**|**Difficulty**|
|-------|-----------|------------|--------------|

`

const main = `
package main

import "fmt"

func main() {
	fmt.Println("test your solution here")
}
`

const doc = `
/*
Package solution solves the leetcode problem 'Two Sum ★'
Category:     algorithms
ID:           1
Title:        Two Sum ★
URL:          https://leetcode.com/problems/two-sum/description/
Difficulty:   Easy (43.17%)
Accepted:     1.7M
Submissions:  3.9M
Test Example: '[2,7,11,15]\n9'
Test Example: '[2,7,11,15]\n9'
Test Example: '[2,7,11,15]\n9'
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
Example:
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
`
