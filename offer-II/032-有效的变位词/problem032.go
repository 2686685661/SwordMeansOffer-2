/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/30 4:11 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _32_有效的变位词


func isAnagram(s string, t string) bool {
	if s == t || len(s) != len(t) {
		return false
	}

	var sa, ta [26]int

	for _, c := range s {
		sa[c - 'a']++
	}
	for _, c := range t {
		ta[c - 'a']++
	}

	return sa == ta
}

func isAnagram2(s string, t string) bool {
	if s == t || len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)

	for i, c := range s {
		m[c]++
		m[rune(t[i])]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}