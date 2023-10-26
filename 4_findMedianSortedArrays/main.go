package main

// PROJECT_NAME:Arithmetic
// DATE:2023/10/26 10:52
// USER:21005

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		nums1, nums2, m, n = nums2, nums1, n, m
	}

	imin, imax, halfLen := 0, m, (m+n+1)/2
	for imin <= imax {
		i := (imin + imax) / 2
		j := halfLen - i
		if i < m && nums2[j-1] > nums1[i] {
			// i 太小，增大 i
			imin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			// i 太大，减小 i
			imax = i - 1
		} else {
			// i 刚好合适
			var maxOfLeft int
			if i == 0 {
				maxOfLeft = nums2[j-1]
			} else if j == 0 {
				maxOfLeft = nums1[i-1]
			} else {
				maxOfLeft = max(nums1[i-1], nums2[j-1])
			}

			if (m+n)%2 == 1 {
				// 总元素个数为奇数，中位数为左半部分的最大值
				return float64(maxOfLeft)
			}

			var minOfRight int
			if i == m {
				minOfRight = nums2[j]
			} else if j == n {
				minOfRight = nums1[i]
			} else {
				minOfRight = min(nums1[i], nums2[j])
			}

			// 总元素个数为偶数，中位数为左半部分最大值和右半部分最小值的平均数
			return float64(maxOfLeft+minOfRight) / 2.0
		}
	}

	return 0.0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
