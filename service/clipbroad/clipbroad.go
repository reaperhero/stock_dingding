package clipbroad

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"regexp"
	"strings"
)

var (
	filePath = "./service/clipbroad/clipbroad.txt"
)

func GetClipBroadRows() ([][]string, error) {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0664)
	if err != nil {
		log.Errorf("Open file error! %v", err)
		return nil, err
	}

	defer func() {
		_ = file.Close()
	}()

	_, err = file.Stat()
	if err != nil {
		return nil, err
	}
	var (
		buf       = bufio.NewReader(file)
		rows      = make([][]string, 0, 3000)
		lineCount = 0
	)
	for {
		line, err := buf.ReadString('\n')
		if lineCount == 0 {
			lineCount++
			continue
		}

		if err == io.EOF {
			return rows, nil
		}
		if err != nil {
			log.Errorf("[GetClipBroadRows] buf.ReadString %v", err)
			return nil, err
		}
		columns := strings.Fields(line)
		for i, column := range columns {
			if column == "" {
				columns = append(columns[:i], columns[i+1:]...)
			}
		}
		stockNameLen := len(columns) - 42
		if stockNameLen > 0 {
			columns[2] = strings.Join(columns[2:stockNameLen+3], "")
			columns = append(columns[:3], columns[stockNameLen+3:]...)
		}
		rows = append(rows, columns)
	}
}

var (
	nicknamePattern = `^[a-zA-Z\p{Han}]+(_[a-zA-Z\p{Han}]+)*?$` // // 用户昵称的正则匹配, 合法的字符有  A-Z, a-z, _, 汉字
	nicknameRegexp  = regexp.MustCompile(nicknamePattern)
)

func isStockNameString(s string) bool {
	if len(s) == 0 {
		return false
	}
	return nicknameRegexp.MatchString(s)
}
