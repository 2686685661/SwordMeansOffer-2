/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/9 4:27 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/


package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("./test.txt")
	if nil != err {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)

	helper := func(s string) string {
		var (
			marks int           // 引号计数
			stmp string         // 中间缓存
			out []string        // 输出
			sRune = []rune(s)
		)

		for _, v := range sRune {
			// 当前字节是逗号且前面没有引号时，添加缓存至输出
			if v == ',' && marks == 0 {
				out = append(out, stmp)
				stmp = ""
				continue
			}
			// 当前字节是引号
			if v == '"' {
				// 缓存为空，引号出现在段落头
				if stmp == "" {
					stmp += "@"
					marks += 1
				} else {
					// 前面已经出现一个引号，转译为真引号
					if stmp[len(stmp) - 1] == '#' {
						t := []rune(stmp)
						stmp = string(append(t[:len(t)-1], '"'))
						marks -= 1
					} else {
						stmp = stmp + "#"
						marks += 1
					}
				}
				continue
			}
			// 当前是逗号，且前面只有一对引号，去除头尾的标识
			if v == ',' && marks == 2 && stmp[len(stmp)-1] == '#' {
				out = append(out, stmp[1:len(stmp)-1])
				stmp, marks = "", 0
			} else {
				// 其他字符，写入缓存
				stmp = string(append([]rune(stmp), v))
			}
		}
		if stmp != "" && stmp[0] == '@' { //当最后一个单词头出现引号时，需去掉标识字符
			out = append(out, stmp[1:len(stmp)-1])
		} else {
			out = append(out, stmp)
		}
		return strings.Join(out, "\t")
	}

	// 主流程
	for {
		line, err := rd.ReadString('\n')
		if nil != err || io.EOF == err {
			break
		}
		line = strings.Trim(line, "\n")
		fmt.Println(helper(line))
	}
}