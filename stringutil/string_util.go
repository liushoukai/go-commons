package stringutil

import "strings"

// IsBlank 判断字符串是否为空或仅包含空白字符
func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsNotBlank 判断字符串是否非空白
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

// IsAnyBlank 判断任意一个字符串是否为空白
func IsAnyBlank(strs ...string) bool {
	for _, s := range strs {
		if IsBlank(s) {
			return true
		}
	}
	return false
}

// IsAllBlank 判断是否所有字符串都为空白
func IsAllBlank(strs ...string) bool {
	for _, s := range strs {
		if IsNotBlank(s) {
			return false
		}
	}
	return true
}

// TrimAllSpace 去除所有字符串的首尾空白
func TrimAllSpace(strs ...string) []string {
	res := make([]string, len(strs))
	for i, s := range strs {
		res[i] = strings.TrimSpace(s)
	}
	return res
}
