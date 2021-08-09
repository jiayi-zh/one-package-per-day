// Author:  jy-zh
// Version: 1.0.0
// Date:    2021/8/9 16:39
// Description:
// 全局变量的初始化按照依赖关系初始化, 优先处理被依赖的, 可以定义多个 init()
// 局部变量的初始化按照定义顺序初始化

package native

import (
	"fmt"
	"testing"
)

func Test_init(t *testing.T) {
	fmt.Println("3. main() 函数初始化")
}
