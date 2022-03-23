/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/29 12:11 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _85_生成匹配的括号


/**
回溯法
参考：https://leetcode-cn.com/problems/IDBivT/solution/jian-dan-yi-dong-javac-pythonjs-pei-yang-mmhs/
 */

func generateParenthesis(n int) []string {
	if n < 1 || n > 8 {
		return nil
	}

	var res []string
	var dfs func(cur string, left, right int)
	dfs = func(cur string, left, right int) {
		//当右括号多于左扩号时，此时，已经不可能进行匹配了，比如"())"，所以左括号要多于或者等于右括号
		if left > n || right > n || left < right {
			return
		}
		if len(cur) == 2 * n && left == right {
			res = append(res, cur)
			return
		}
		dfs(cur + "(", left + 1, right)
		dfs(cur + ")", left, right + 1)
	}
	dfs("", 0, 0)
	return res
}