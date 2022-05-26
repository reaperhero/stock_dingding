package spider

import (
	"bytes"
	"encoding/json"
	"github.com/gocolly/colly"
	"github.com/reaperhero/stock_dingding/model"
	log "github.com/sirupsen/logrus"
)

func (s *spider) handleResponse() {
	s.c.OnResponse(func(response *colly.Response) {
		if response.Request.URL.String() == URLBusinessTransaction {
			startIndex := bytes.IndexByte(response.Body, '{')
			lastIndex := bytes.LastIndexByte(response.Body, '}')
			if startIndex <0 ||  lastIndex  <0 {
				log.Errorf("no found")
			}
			data:=model.BusinessTransactionResponse{}
			err := json.Unmarshal(response.Body[startIndex:lastIndex+1],&data)
			if err!=nil{
				log.Errorf("%v",err)
			}
			log.Println(data)
		}
	})
}
