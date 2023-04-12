package main

import "fmt"

func diagonalPrime(nums [][]int) int {
	pre := 0
	for i := 0; i < len(nums); i++ {
		a := nums[i][i]
		b := nums[i][len(nums)-i-1]
		if isPrime(a) && a > pre {
			pre = a
		}
		if isPrime(b) && b > pre {
			pre = b
		}
	}
	return pre
}

func isPrime(num int) bool {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return num >= 2
}

func main() {
	nums := make([][]int, 3)
	nums[0] = []int{1, 2, 3}
	nums[1] = []int{5, 6, 7}
	nums[2] = []int{9, 10, 11}
	fmt.Println(diagonalPrime(nums))
}
