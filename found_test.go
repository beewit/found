package main

import (
	"testing"
	"fmt"
)

func TestRedisList(t *testing.T) {
	var startUrl = make([]string, 2623)
	for i := 0; i < 2623; i++ {
		startUrl[i] = fmt.Sprintf("http://www.weixinqun.com/group?p=%d", i)
	}
	t.Error(startUrl[0])
}
