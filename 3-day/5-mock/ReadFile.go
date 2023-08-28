//package file
//
//import (
//	"bufio"
//	"os"
//	"strings"
//)
//
//func ReadFirstLine() string {
//	open, err := os.Open("log")
//	defer open.Close()
//	if err != nil {
//		return ""
//	}
//	scanner := bufio.NewScanner(open)
//	for scanner.Scan() { // 逐行扫描文件内容
//		return scanner.Text() // 返回第一行文本
//	}
//	return ""
//}
//
//func ProcessFirstLine() string {
//	line := ReadFirstLine() // 第一行文本在文件中是:line11
//	destLine := strings.ReplaceAll(line, "11", "00")
//	return destLine
//}

package file

import (
	"bufio"
	"os"
	"strings"
)

type FileReader interface {
	ReadFirstLine() string
}

type FileReaderImpl struct{}

// 需要插桩
func (f *FileReaderImpl) ReadFirstLine() string {

	open, err := os.Open("log")
	defer open.Close()
	if err != nil {
		return ""
	}
	scanner := bufio.NewScanner(open)
	for scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func ProcessFirstLine(reader FileReader) string {
	line := reader.ReadFirstLine()
	destLine := strings.ReplaceAll(line, "11", "00")
	return destLine
}
