/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/19 8:40 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _16_数值的整数次方

func Power( base float64 ,  exponent int ) float64 {
	if base == 0.0 && exponent < 0 {
		return 0
	}

	uexp := exponent
	if exponent < 0 {
		uexp = 0 - exponent
	}

	result := powerCore(base, uexp)

	if exponent < 0 {
		result = 1.0 / result
	}
	return result
}

func powerCore(base float64, exponent int) float64 {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}

	result := powerCore(base, exponent >> 1)
	result *= result
	if exponent & 0x1 == 1 {
		result *= base
	}
	return result
}