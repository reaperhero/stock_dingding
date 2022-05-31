package dingding

import "github.com/blinkbean/dingtalk"

var (
	DingClient dingClient
)

type dingClient struct {
	client *dingtalk.DingTalk
}
//
//func init() {
//	// 单个机器人有单位时间内消息条数的限制，如果有需要可以初始化多个token，发消息时随机发给其中一个机器人。
//	var dingToken = []string{"7bd675b66646ba890046c2198257576470099e1bda0770bad7dd6684fb1e0415"}
//	cli := dingtalk.InitDingTalk(dingToken, ".")
//	cli.SendTextMessage("content")
//}

func (d *dingClient) SendStockPriceRanking() {

}
