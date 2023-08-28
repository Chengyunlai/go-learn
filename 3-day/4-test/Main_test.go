package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain(m *testing.M) {
	println("测试之前")
	m.Run()
	println("测试之后")
}

func TestSearchNameById(t *testing.T) {
	output := searchNameById(1)
	expectOutPut := "Admin"
	//if output != expectOutPut {
	//	t.Errorf("预期值是 %s，返回值是 %s", expectOutPut, output)
	//}
	assert.Equal(t, expectOutPut, output)
}
