package homework01

import (
	"strconv"
)

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	var result = 0
	for _, v := range nums {
		result = result ^ v
	}
	return result
}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	s := strconv.Itoa(x)
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

// 3. 有效的括号
// 给定一个只包括 '(', , '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool {
	var runes = []byte(s)
	stack := make([]byte, 0)

	for _, v := range runes {
		switch v {
		case 40, 123, 91:
			stack = append(stack, v)
		case 41:
			if len(stack) != 0 && stack[len(stack)-1] == 40 {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, v)
			}
		case 125:
			if len(stack) != 0 && stack[len(stack)-1] == 123 {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, v)
			}
		case 93:
			if len(stack) != 0 && stack[len(stack)-1] == 91 {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, v)
			}
		}
	}
	return len(stack) == 0
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	// TODO: implement
	return ""
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	i := len(digits) - 1
	for ; i >= 0; i-- {
		if digits[i]+1 != 10 {
			digits[i] = digits[i] + 1
			return digits
		} else {
			digits[i] = 0
		}
	}
	result := make([]int, len(digits)+1)
	result[0] = 1
	for j := 1; j < len(result); j++ {
		result[j] = digits[j-1]
	}
	return result
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	for i := 1; i < len(nums); {
		if nums[i] == nums[i-1] {
			i++
		} else {
			i++
		}
	}
	return 0
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	// TODO: implement
	return nil
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		value, ok := m[target-v]
		if ok {
			return []int{i, value}
		} else {
			m[v] = i
		}
	}
	return nil
}
