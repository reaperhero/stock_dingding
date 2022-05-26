package spider

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	log "github.com/sirupsen/logrus"
)

type spider struct {
	c *colly.Collector
}

func NewSpider() *spider {
	c := colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{}),
		//colly.AllowedDomains(_domains...),
		colly.MaxDepth(0),
	)
	s := &spider{c}
	s.setSpider()
	return s
}

func (s *spider) setSpider() {
	s.c.OnError(func(response *colly.Response, err error) {
		log.Errorf("Request URL: %s failed with response: %s err %v\n", response.Request.URL, string(response.Body), err)
	})
	s.c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	s.handleResponse()
}


func (s *spider) BusinessTransaction() {
	err := s.c.Visit(URLBusinessTransaction)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
}
