package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

// 基本思路:
// 编码时用4字节记录str长度，形如 [len1]str1[len2]str2...
type Codec struct {
}

func (c *Codec) Encode(strs []string) string {
}

func (c *Codec) Decode(strs string) []string {
}

func main() {
	var codec Codec
	fmt.Println(codec.Decode(codec.Encode([]string{"hello", "go", "!"})))
}
