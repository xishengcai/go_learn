package stringset

func maxArea(height []int) int {
	l, r := 0, len(height)-1

	max := (r - l) * min(height[l], height[r])

	for l < r {
		if height[l] <= height[r] {
			l++
			if height[l] < height[l-1] {
				continue
			}

			if min(height[l], height[r])*(r-l) > max {
				max = min(height[l], height[r]) * (r - l)
			}
		} else {
			r--
			if height[r] < height[r+1] {
				continue
			}

			if min(height[l], height[r])*(r-l) > max {
				max = min(height[l], height[r]) * (r - l)
			}
		}
	}
	return max
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}
