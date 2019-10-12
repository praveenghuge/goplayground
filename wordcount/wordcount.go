package wordcount

import (
	"fmt"
	"strings"
)

// Word Count ...
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	newS := strings.Split(s, " ")
	for i := 0; i < len(newS); i++ {
		_, ok := m[newS[i]]
		m[newS[i]] = 1
		if ok == true {
			m[newS[i]]++
		}
	}
	fmt.Println(m)
	return m

}
