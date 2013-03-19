/**
 * Created with IntelliJ IDEA.
 * User: lucas
 * Date: 3/19/13
 * Time: 2:20 PM
 * To change this template use File | Settings | File Templates.
 */
package ex40

import (
	"wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	fields := strings.Fields(s)

	for _, v := range fields {
		if _, ok := m[v]; ok {
			m[v]++
		}else {
			m[v] = 1
		}
	}

	return m
}

func Main() {
	wc.Test(WordCount)
}
