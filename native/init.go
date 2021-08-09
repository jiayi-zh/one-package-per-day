// Author:  jy-zh
// Version: 1.0.0
// Date:    2021/8/9 16:36
// Description: Go初始化
//

package native

import "fmt"

var p int64 = a()

func a() int64 {
	fmt.Println("1. 变量函数初始化")
	return 0
}

func init() {
	fmt.Println("2. init() 函数初始化")
}
