package bucket

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 方法一： 暴力循环
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if abs(nums[i])-(nums[j]) <= t && abs(i-j) <= k {
				println(i, " ", j)
				return true
			}
		}
	}
	return false
}

// 划分桶
func getID(x, w int) int {
	if x >= 0 {
		return x / w
	}
	return (x+1)/w - 1
}

// 方法二： 桶排序
// item/t+1 都在一个桶内， 最大和最小差值正好是t
func containsNearbyAlmostDuplicate3(nums []int, k, t int) bool {
	bucket := map[int]int{}
	for i, x := range nums {

		id := getID(x, t+1)
		// 在同一个桶内的数据之差 <= t
		if _, has := bucket[id]; has {
			return true
		}

		if y, has := bucket[id-1]; has && abs(x-y) < t {
			return true
		}

		if y, has := bucket[id+1]; has && abs(x-y) < t {
			return true
		}

		bucket[id] = x
		if i >= k {
			// 滑动窗口，最左侧移除
			delete(bucket, getID(nums[i-k], t+1))
		}
	}
	return false
}

func containsNearbyAlmostDuplicate4(nums []int, k, t int) bool {
	mp := map[int]int{}
	for i, x := range nums {
		id := getID(x, t+1)
		if _, has := mp[id]; has {
			return true
		}
		if y, has := mp[id-1]; has && abs(x-y) <= t {
			return true
		}
		if y, has := mp[id+1]; has && abs(x-y) <= t {
			return true
		}
		mp[id] = x
		if i >= k {
			delete(mp, getID(nums[i-k], t+1))
		}
	}
	return false
}
