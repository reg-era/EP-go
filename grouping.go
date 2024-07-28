package ep

import (
	"fmt"
	"strings"
)

func Grouping(args []string) string {
	res, patr := "", ""
	find := []string{}

	// get patern
	if len(args[0]) > 2 && args[0][0] == '(' && args[0][len(args[0])-1] == ')' {
		patr = args[0][1 : len(args[0])-1]
	} else {
		return ""
	}

	// find match
	sls := strings.Split(args[1], " ")
	slp := strings.Split(patr, "|")
	for i := 0; i < len(sls); i++ {
		for j := 0; j < len(slp); j++ {
			if strings.Contains(sls[i], slp[j]) {
				find = append(find, sls[i])
			}
		}
	}

	// check ponctuation
	for i := 0; i < len(find); i++ {
		for j := len(find[i]) - 1; j >= 0; j-- {
			if find[i][j] == ',' {
				find[i] = find[i][:j]
			}
		}
	}

	// make respons
	for i := 0; i < len(find); i++ {
		res += fmt.Sprintf("%d: %s", i+1, find[i]) + "\n"
	}
	return res
}
