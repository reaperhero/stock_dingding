package clipbroad

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
	"time"
	"unicode"
)

var (
	filePath = "./service/clipbroad/clipbroad.txt"
)

func GetClipBroadRows() ([][]string, error) {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err != nil {
		log.Errorf("Open file error! %v", err)
		return nil, err
	}
	defer file.Close()

	_, err = file.Stat()
	if err != nil {
		return nil, err
	}
	var (
		buf  = bufio.NewReader(file)
		rows = make([][]string, 0, 3000)
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
		var (
			lineRow  = make([]string, 0, 42)
			previous = ""
		)
		continueFor:
		for _, s := range strings.Split(line, "   ") {
			column := strings.TrimSpace(s)
			if column == "" {
				continue continueFor
			}
			switch previous {
			case "":
				for _, r := range column {
					if unicode.Is(unicode.Han, r) {
						previous = column
						lineRow = append(lineRow, column)
						continue continueFor
					}
				}
			default:
				if unicode.IsNumber(rune(column[0])){
					previous = ""
					lineRow = append(lineRow, column)
					continue continueFor
				}
				for _, r := range column {
					if unicode.Is(unicode.Han, r) {
						previous = column
						lineRow[len(lineRow)-1] = lineRow[len(lineRow)-1] + column
						continue continueFor
					}
				}

			}
		}
		previous = ""
		if len(lineRow) == 42 {
			rows = append(rows, lineRow)
			continue
		}
		for _, s := range lineRow {
			fmt.Println(s)
		}
		time.Sleep(time.Second * 12)
	}
}
