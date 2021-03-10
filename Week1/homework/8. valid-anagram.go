package homework

/*
@Time    : 2021/3/10 20:26
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 8. valid-anagram.go
@Software: GoLand
*/

/*
242. 有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
说明:
你可以假设字符串只包含小写字母
*/

func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// use bitMap record element count
	var bitMap [26]int
	for _, v := range []rune(s) {
		bitMap[v-'a']++
	}
	for _, t := range []rune(t) {
		bitMap[t-'a']--
	}
	for i := range bitMap {
		if bitMap[i] != 0 {
			return false
		}
	}
	return true
}
