package main

import (
	"strings"
)

func findTheLongestBalancedSubstring(s string) int {
	if !strings.ContainsAny(s, "0") || !strings.Contains(s, "1") {
		return 0
	}
	a := "01"
	for strings.Contains(s, a) {
		a = "0" + a + "1"
	}
	return len(a) - 2
}
