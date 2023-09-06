package pkg

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ReadFileStringLine(path string) (err error, pathList []string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err || line == "\n" {
			if line == "" || line == "\n" {
				break
			}
		}
		if line[0] == 0xef || line[1] == 0xbb || line[2] == 0xbf {
			line = line[3:]
		}
		line = strings.Replace(line, "\n", "", -1)
		line = strings.Replace(line, "\r\n", "", -1)
		if line != "" {
			pathList = append(pathList, line)
		}
	}
	return nil, pathList
}

func IsExist(subStr string, strs []string) bool {
	for _, single := range strs {
		if subStr == single {
			return true
		}
	}
	return false
}
func SubIsContains(strList []string, str string) (bool, []int) {
	isContains := false
	var sub []int
	for index, value := range strList {
		if strings.Contains(value, str) {
			isContains = true
			sub = append(sub, index)
		}
	}
	return isContains, sub
}
