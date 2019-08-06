package test

import "testing"

func TestMD5(t *testing.T) {
	md5 := MD5("123456")
	if md5 != "e10adc3949ba59abbe56e057f20f883e" {
		t.Error("测试失败")
	}
}
