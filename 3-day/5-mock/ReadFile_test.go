package file

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	file "top.chengyunlai/go-learn/3-day/5-mock/mocks"
)

func TestMain(m *testing.M) {
	println("测试之前")
	m.Run()
	println("测试之后")
}

// 没有Mock
func TestProcessFirstLine(t *testing.T) {
	mockReader := &FileReaderImpl{}      // 创建 MockFileReader 对象
	line := ProcessFirstLine(mockReader) // 调用原来的方法
	expected := "line00"                 // 预期的返回值
	assert.Equal(t, expected, line)
}

// 有Mock
func TestProcessFirstLineWithMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// 注意，主要是这里，file是包名。
	mockFileReader := file.NewMockFileReader(ctrl)
	mockFileReader.EXPECT().ReadFirstLine().Return("line11")

	result := ProcessFirstLine(mockFileReader)
	expected := "line00"

	if !strings.EqualFold(result, expected) {
		t.Errorf("ProcessFirstLine returned unexpected result. Got: %s, Expected: %s", result, expected)
	}
}
