/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/23 8:45 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/


package main

import (
	"fmt"
	"sync"
)

func main() {
	a := []int{1, 1, 1}
	b := []int{2, 2, 2}
	c := []int{3, 3, 3}
	fmt.Println(calculate(a, b, c))
}

func calculate(in ...[]int) []int {
	var result []int = make([]int, 3)
	wg := sync.WaitGroup{}
	wg.Add(len(in))

	for i, s := range in {
		sum := 0
		go func(index int, arr []int) {
			for _, v := range arr {
				sum += v
			}
			result[index] = sum
			//result = append(result, sum)
			wg.Done()
		}(i, s)
	}
	wg.Wait()
	return result
}



// 配送-商品管理
// go + mysql + redis + kafka + aws

// merchant_table: id，name, score,地点
// id primary_key
// name prefix(name, n)
// name_score_idx key(name, score)

// mer_ch_table: 优惠，id,mid，限额，优惠，使用时间，状态
// promary_key cid
// mid_cid_idx key(mid, cid)

// mer_food_table: foodid,mid, ... 价格，状态，制作时间, food_type
// primary_key foodid
// mid_fid_idx key(mid, fid)
// mem_idx key(mem)
// type_idx key(fooed_type)

// food_type_table: 菜品类型 id fooid,



// getMerchants(name, sore，距离，菜品。。。)
// 获取商家优惠
// 查询商家菜品（）





