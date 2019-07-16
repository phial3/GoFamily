// 二分法的四种变种
package main

func ER1(x []int, value int) int {
	low := 0
	heigh := len(x) - 1
	for low <= heigh {
		middle := low + (heigh-low)>>1 // 这样为了防止 数字过大。
		if x[middle] == value {
			if middle == 0 || x[middle-1] != value { // 如果前后都是相同的了，那么肯定就是了，再有 前面那个不是 后面那个是 肯定就是符合了，
				// else 说明 前面的是 后面的也是，那么就不对，所以要取前面的那个区间，因为第一个在前面那个区间内
				return middle
			} else {
				heigh = middle - 1
			}
		} else if x[middle] > value {
			heigh = middle - 1
		} else {
			low = middle + 1
		}
	}
	return -1
}
