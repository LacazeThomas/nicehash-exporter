package models

type Feeinfo struct {
	Coin      string `json:"coin"`
	Intervals []struct {
		Start   int     `json:"start"`
		End     float64 `json:"end"`
		Element struct {
			Value    float64     `json:"value"`
			Type     string      `json:"type"`
			Sndvalue interface{} `json:"sndValue"`
			Sndtype  interface{} `json:"sndType"`
		} `json:"element"`
	} `json:"intervals"`
}

type Feerules struct {
	Withdrawal struct {
		External struct {
			Rules struct {
				Btc struct {
					Coin      string `json:"coin"`
					Intervals []struct {
						Start   float64 `json:"start"`
						End     float64 `json:"end"`
						Element struct {
							Value    float64     `json:"value"`
							Type     string      `json:"type"`
							Sndvalue interface{} `json:"sndValue"`
							Sndtype  interface{} `json:"sndType"`
						} `json:"element"`
					} `json:"intervals"`
				} `json:"BTC"`
			} `json:"rules"`
		} `json:"EXTERNAL"`
	}
}
