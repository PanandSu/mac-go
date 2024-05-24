package main

import (
	"fmt"
	"strings"
)

type (
	solution1 struct{}
	solution2 struct{}
	solution3 struct{}
)

func main() {
	word := "transform"
	wordDict := []string{"trans", "tran", "sf", "orm", "for"}
	fmt.Println(solution1{}.canSegmentWord(word, wordDict))
	fmt.Println(solution2{}.canSegmentWord(word, wordDict))
	fmt.Println(solution3{}.wordBreak(word, wordDict))
}

func (s1 solution1) canSegmentWord(word string, wordDict []string) bool {
	if word == "" {
		return true
	}
	for _, w := range wordDict {
		if strings.HasPrefix(word, w) {
			if s1.canSegmentWord(word[len(w):], wordDict) {
				return true
			}
		}
	}
	return false
}

func (s2 solution2) canSegmentWord(word string, wordDict []string) bool {
	m := make(map[string]bool)
	return s2.canSegmentHelper(word, wordDict, m)
}

func (s2 solution2) canSegmentHelper(word string, wordDict []string, mem map[string]bool) bool {
	if word == "" {
		return true
	}
	// 先去备忘录看看是否有记录
	if result, ok := mem[word]; ok {
		return result
	}
	for _, w := range wordDict {
		if strings.HasSuffix(word, w) {
			if s2.canSegmentHelper(word[len(w):], wordDict, mem) {
				mem[word] = true //记录缓存值
				return true
			}
		}
	}
	mem[word] = false //记录缓存值
	return false
}

// dp[0]=true
// dp[1]=false
// dp[2]=false
// dp[3]=false
// dp[4]=true
// dp[5]=true

func (s3 solution3) wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool)
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for _, word := range wordDict {
		m[word] = true
	}
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && m[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
