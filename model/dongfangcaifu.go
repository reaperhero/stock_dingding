package model

type BusinessTransactionResponse struct {
	Result struct {
		Data []BusinessTransactionStock `json:"data"`
	} `json:"result"`
}

type BusinessTransactionStock struct {
	OPERATEDEPTCODE    string      `json:"OPERATEDEPT_CODE"`
	OPERATEDEPTNAME    string      `json:"OPERATEDEPT_NAME"`
	TRADEDATE          string      `json:"TRADE_DATE"`
	D1CLOSEADJCHRATE   interface{} `json:"D1_CLOSE_ADJCHRATE"`
	D2CLOSEADJCHRATE   interface{} `json:"D2_CLOSE_ADJCHRATE"`
	D3CLOSEADJCHRATE   interface{} `json:"D3_CLOSE_ADJCHRATE"`
	D5CLOSEADJCHRATE   interface{} `json:"D5_CLOSE_ADJCHRATE"`
	D10CLOSEADJCHRATE  interface{} `json:"D10_CLOSE_ADJCHRATE"`
	SECURITYCODE       string      `json:"SECURITY_CODE"`
	SECURITYNAMEABBR   string      `json:"SECURITY_NAME_ABBR"`
	ACTBUY             float32         `json:"ACT_BUY"`
	ACTSELL            float32         `json:"ACT_SELL"`
	NETAMT             float32         `json:"NET_AMT"`
	EXPLANATION        string      `json:"EXPLANATION"`
	D20CLOSEADJCHRATE  interface{} `json:"D20_CLOSE_ADJCHRATE"`
	D30CLOSEADJCHRATE  interface{} `json:"D30_CLOSE_ADJCHRATE"`
	SECUCODE           string      `json:"SECUCODE"`
	OPERATEDEPTCODEOLD string      `json:"OPERATEDEPT_CODE_OLD"`
	ORGNAMEABBR        string      `json:"ORG_NAME_ABBR"`
	CHANGERATE         float64     `json:"CHANGE_RATE"`
}
