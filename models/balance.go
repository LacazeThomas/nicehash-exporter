package models

type Balance struct {
	Available      string `json:"available"`
	Totalbalance   string `json:"totalBalance"`
	Pendingdetails struct {
		Unpaidmining    string `json:"unpaidMining"`
		Hashpowerorders string `json:"hashpowerOrders"`
		Exchange        string `json:"exchange"`
		Deposit         string `json:"deposit"`
		Withdrawal      string `json:"withdrawal"`
	} `json:"pendingDetails"`
	Btcrate  float32 `json:"btcRate"`
	Currency string  `json:"currency"`
	Active   bool    `json:"active"`
	Pending  string  `json:"pending"`
}
