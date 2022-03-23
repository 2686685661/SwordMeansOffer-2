/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/17 1:45 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _63_买卖股票的最好时机_一_


func maxProfit( prices []int ) int {

	if prices == nil || len(prices) < 2 {
		return 0
	}

	min := prices[0]
	maxProfit := prices[1] - min

	for i := 2; i < len(prices); i++ {
		if prices[i - 1] < min {
			min  =prices[i - 1]
		}
		current := prices[i] - min
		if current  > maxProfit {
			maxProfit = current
		}
	}
	if maxProfit < 0 {
		return 0
	}
	return maxProfit
}