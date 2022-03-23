/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/30 5:22 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _33_变位词组

import "sort"

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	m := make(map[string][]string)

	for _, s := range strs {
		bs := []byte(s)
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		m[string(bs)] = append(m[string(bs)], s)
	}

	for _, v := range m {
		res = append(res, v)
	}
	return res
}