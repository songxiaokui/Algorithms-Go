package homework

import "sort"

/*
@Time    : 2021/3/11 12:59
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 12.group anagram.go
@Software: GoLand
*/

/*
49. 字母异位词分组
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
说明：

所有输入均为小写字母。
不考虑答案输出的顺序

*/

func GroupAnagram(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	// use hashmap save sort key value
	hashMap := make(map[string][]string)

	for _, v := range strs {
		s2rune := []rune(v)
		// min sort
		sort.Slice(s2rune, func(i, j int) bool {
			return s2rune[i] < s2rune[j]
		})
		// convert string
		rune2string := string(s2rune)
		// add to hashMap
		hashMap[rune2string] = append(hashMap[rune2string], v)
	}

	result := make([][]string, 0, len(hashMap))
	for _, v := range hashMap {
		result = append(result, v)
	}
	return result
}
