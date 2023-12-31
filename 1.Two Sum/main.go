// https://leetcode.com/problems/two-sum/description/
package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0,1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // [1,2]
	fmt.Println(twoSum([]int{3, 3}, 6))         // [0,1]
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ { //внешний цикл
		for j := i + 1; j < len(nums); j++ { //внутренний цикл
			if target == nums[i]+nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

