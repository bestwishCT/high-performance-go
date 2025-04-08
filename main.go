package main

import (
	"fmt"
	"strings"
)

func main() {
	var mySlogaj strings.Builder
	mySlogaj.WriteString("Go")
	mySlogaj.WriteString("lang")
	mySlogaj.WriteByte('!')        // 追加单个字符
	fmt.Println(mySlogaj.String()) // 输出：Golang!
}
