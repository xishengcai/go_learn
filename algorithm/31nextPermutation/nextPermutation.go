package _1nextPermutation

// 题解
// 1. 找到 逆序的
// 2. 将 比他大的调换位置
// 3. 倒叙排列

func nextPermutation(nums []int) {
	l := len(nums)

	i := l - 2

	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}

	for j := l - 1; j > 0 && i > -1; j-- {
		if nums[j] > nums[i] {
			nums[i], nums[j] = nums[j], nums[i]
			break
		}
	}

	reverse(nums[i+1:])
	for i := 0; i < l; i++ {
		print(nums[i], " ")
	}
}

func reverse(nums []int) {
	l := len(nums)
	i := 0
	n := l - 1
	for i < n {
		nums[i], nums[n] = nums[n], nums[i]
		i++
		n--
	}
}
