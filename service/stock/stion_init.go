package stock

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
)

func RecordFileWithStions(ss []Siton) {
	file, _ := os.OpenFile(fileRdName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer func() {
		_ = file.Close()
	}()

	marshal, err := json.MarshalIndent(ss, "", " ")
	if err != nil {
		log.Errorf("[SitonManag] RecordFile json.Marshal %v", err)
		return
	}
	file.Write(marshal)
}