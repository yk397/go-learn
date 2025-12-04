package main

import (
	"fmt"
	"sort"
	"strconv"
)

//只出现一次的数字
/*
给定一个非空整数数组，除了某个元素只出现一次以外，
其余每个元素均出现两次。找出那个只出现了一次的元素。
可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/

func SingleNumber(nums []int) {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v == 1 {
			fmt.Println(k)
		}
	}
}

//回文数
/*
考察：数字操作、条件判断
题目：判断一个整数是否是回文数
*/
func IsPalindrome(x int) bool {
	str := strconv.Itoa(x)
	var newstr string
	for i := len(str) - 1; i >= 0; i-- {
		newstr = fmt.Sprintf("%s%c", newstr, str[i])
	}
	return str == newstr
}

//有效括号
/*
考察：字符串处理、栈的使用

题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，
判断字符串是否有效
*/
func IsValidBracket(s string) bool {
	stack := []rune{}
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if (c == ')' && top == '(') || (c == '}' && top == '{') || (c == ']' && top == '[') {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}

//最长公共前缀
/*
考察：字符串处理、循环嵌套

题目：查找字符串数组中的最长公共前缀
*/

func LongestCommonPrefix(strs []string) string {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	if len(strs) == 0 {
		return ""
	}
	for i, c := range strs[0] {
		for _, e := range strs {
			if rune(e[i]) != c {
				return strs[0][:i]
			}
		}
	}
	return ""
}

//加一
/*
难度：简单

考察：数组操作、进位处理

题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
*/

func AddOne(nums []int) []int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < 9 {
			nums[i]++
			break
		}
		if i == len(nums)-1 {
			nums[i] = 0
		} else {
			nums[i] = 1
		}
		if i == 0 {
			nums = append([]int{1}, nums...)
		}
	}
	return nums
}

//删除有序数组中的重复项
/*
给你一个有序数组 nums ，请你原地删除重复出现的元素，使
每个元素只出现一次，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并
在使用 O(1) 额外空间的条件下完成。可以使用双指针法，
一个慢指针 i 用于记录不重复元素的位置，一个快指针 j
用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j]
赋值给 nums[i + 1]，并将 i 后移一位。
*/

func RemoveDuplicates(nums []int) int {
	repeat := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[repeat] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		} else {
			repeat = i
		}
	}
	return len(nums)
}

//合并区间
/*
以数组 intervals 表示若干个区间的集合，
其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，
该数组需恰好覆盖输入中的所有区间。可以先对区间数组按照区
间的起始位置进行排序，然后使用一个切片来存储合并后的区间，
遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，
如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中
*/

func Merge(intervals [][]int) [][]int {
	results := [][]int{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	temp := intervals[0]
	if len(intervals) > 1 {
		for i := 1; i < len(intervals); i++ {
			if intervals[i][0] > temp[1] {
				results = append(results, temp)
				temp = intervals[i]
			} else {
				temp[1] = intervals[i][1]
			}
		}
		results = append(results, temp)
	}
	return results
}

//两数之和
/*
考察：数组遍历、map使用

题目：给定一个整数数组 nums 和一个目标值
 target，请你在该数组中找出和为目标值的那两个整数
*/

func TowSum(nums []int, target int) [2]int {
	result := [2]int{0}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; i < len(nums)-1; i++ {
			if nums[i]+nums[j] == target {
				return [2]int{nums[i], nums[j]}
			}
		}
	}

	return result
}
