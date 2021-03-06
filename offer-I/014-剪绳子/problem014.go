/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/19 2:30 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _14_剪绳子

// 动态规划
/*
首先，我们知道如果一个长度为n的绳子可以分为 n = s[1] + s[2] + s[3] + s[4] 这么几个段，
那么一定有 max = s[1]*s[2]*s[3]*s[4], 且 s[1] * s[2] 为长度是 s[1]+[s2]的子问题的最优解,同理s[3]*s[4]一定是s[3]+s[4]的最优解,所以我们知道 n 其实可以表示为上述两段的最优解.

我想这样你就对本问题有了一定的认识:
定义一个数组 max[0...61] 这里我们定义为长度为i(下标)的绳子能够构成的最优解长度,那么我们只要把这个定义成为一个二分问题，就是最优解的二分问题,（原理是:本问题的最优解一定是由子问题的最优解构成,如果你不能明白,那么请回过去看第一段话.）如果我们从最小的开始打表,那么试一下每两个问题的最优解，找到其中最大值就能得到本问题的解。
那么我们能够得到递推公式:
max[i] = Max(max[1]*max[i-1],max[2]*max[i-2], .... , max[i-1]*max[1], i )  当 i ≠ target时
max[i] = Max(max[1]*max[i-1],max[2]*max[i-2], .... , max[i-1]*max[1])     当 i ＝ target时
这里你可能会有疑问，为什么我要包括 i 当不是目标解的时候, 也要试i,试这样的，我们继续看一下题干:
---------
给你一根长度为n的绳子，请把绳子剪成整数长的m段（m、n都是整数，n>1并且m>1），每段绳子的长度记为k[0],k[1],...,k[m]。请问k[0]xk[1]x...xk[m]可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。
-------- 发现绳子最少要切一刀,这个是一个比较烦的点, 首先举个例子。
f[长度] = 最少切一刀的最优解; 最少切一刀或者不切的最优解
f[1] = 1; 1 f[2] = 1, 2 f[3] = 2, 3 f[4] = 4, 4 f[5] = 6, 6 f[6] = Max( f[1]*1[5],f[2]*f[4] f[3] * [3]) = 9
这是我在计算f[6]时的过程，如果计算f[6],我们需要计算这些子问题的最优解
f[6] = Max( f[1]*1[5],f[2]*f[4] f[3] * [3]) = 9
那么可以看出来,我们对 6 做最优的二分已经保证了最少切1次,对于
f[6] = f[ x ] * f[6-x] 这两个子部分，最优解并不是必须切1次，可以不切.
f[6] = f[3] * f[3] 从结果上我们知道是3,那么就是3 * 3, 但是f[3] 如果二分结果并不是3 而是2最好的分法是不分, f[3] = 3, 这样你就理解了我为什么要加上 i 做最大值的比较，也理解了为什么不用i 做最终target的比较了吧。
*/
func CutRope( number int ) int {
	if number >0 && number <= 3 {
		return number - 1
	}

	products :=make([]int, number + 1)

	products[1] = 1
	for i := 2; i <= number; i++ {
		max := i
		for j := 1; j <= i / 2; j++ {
			if products[j] * products[i-j] > max {
				max = products[j] * products[i-j]
			}
		}
		products[i] = max
	}

	res := -1
	for i := 1; i <= (number / 2); i++ {
		if products[i] * products[number - i] > res {
			res  = products[i] * products[number - i]
		}
	}
	return res
}