/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/28 1:46 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _38_字符串的排列


/**
全排列：https://www.cnblogs.com/cxjchen/p/3932949.html
排列组合公式：https://blog.csdn.net/weixin_30598225/article/details/97546846
 */
func Permutation( str string ) []string {
	if len(str) == 0 {
		return nil
	}
	sch := []byte(str)
	var list []string
	list = permutationCore(sch, 0, list)
	return list
}

func permutationCore(sch []byte, begin int, list []string) []string {
	if begin == len(sch) - 1 {
		list = append(list, string(sch))
	} else {
		set := map[byte]bool{}
		for i := begin; i < len(sch); i++ {
			if _, ok := set[sch[i]]; !ok {
				set[sch[i]] = true
				sch[begin], sch[i] = sch[i], sch[begin]
				list = permutationCore(sch, begin+1, list)
				sch[begin], sch[i] = sch[i], sch[begin]
			}
		}
	}
	return list
}