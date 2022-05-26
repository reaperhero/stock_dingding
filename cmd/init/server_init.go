package init

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat:  	"2006-01-02 15:03:04",
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		DataKey:           "",
		FieldMap:          nil,
		CallerPrettyfier:  nil,
		PrettyPrint:       true,
	})

	// 输出stdout而不是默认的stderr，也可以是一个文件
	//file, err := os.OpenFile(LOG_FILE, os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0755)
	log.SetOutput(os.Stdout)

	// 只记录严重或以上警告
	log.SetLevel(log.DebugLevel)
}
