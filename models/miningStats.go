package models

import (
	"time"
)

type Temperature struct {
	T float64
}
type MiningStats struct {
	Minerstatuses struct {
		Mining int `json:"MINING"`
	} `json:"minerStatuses"`
	Rigtypes struct {
		Managed int `json:"MANAGED"`
	} `json:"rigTypes"`
	Totalrigs          int     `json:"totalRigs"`
	Totalprofitability float64 `json:"totalProfitability"`
	Grouppowermode     string  `json:"groupPowerMode"`
	Totaldevices       int     `json:"totalDevices"`
	Devicesstatuses    struct {
		Mining   int `json:"MINING"`
		Disabled int `json:"DISABLED"`
	} `json:"devicesStatuses"`
	Unpaidamount        string        `json:"unpaidAmount"`
	Path                string        `json:"path"`
	Btcaddress          string        `json:"btcAddress"`
	Nextpayouttimestamp time.Time     `json:"nextPayoutTimestamp"`
	Lastpayouttimestamp time.Time     `json:"lastPayoutTimestamp"`
	Miningriggroups     []interface{} `json:"miningRigGroups"`
	Miningrigs          []struct {
		Rigid            string `json:"rigId"`
		Type             string `json:"type"`
		Name             string `json:"name"`
		Statustime       int64  `json:"statusTime"`
		Jointime         int    `json:"joinTime"`
		Minerstatus      string `json:"minerStatus"`
		Groupname        string `json:"groupName"`
		Unpaidamount     string `json:"unpaidAmount"`
		Softwareversions string `json:"softwareVersions"`
		Devices          []struct {
			ID         string `json:"id"`
			Name       string `json:"name"`
			Devicetype struct {
				Enumname    string `json:"enumName"`
				Description string `json:"description"`
			} `json:"deviceType"`
			Status struct {
				Enumname    string `json:"enumName"`
				Description string `json:"description"`
			} `json:"status"`
			Temperature                    float64 `json:"temperature"`
			VRAMTemperature                int
			GPUTemperature                 int
			Load                           float32 `json:"load"`
			Revolutionsperminute           float32 `json:"revolutionsPerMinute"`
			Revolutionsperminutepercentage float32 `json:"revolutionsPerMinutePercentage"`
			Powermode                      struct {
				Enumname    string `json:"enumName"`
				Description string `json:"description"`
			} `json:"powerMode"`
			Powerusage float32 `json:"powerUsage"`
			Speeds     []struct {
				Algorithm     string `json:"algorithm"`
				DisplaySuffix string `json:"displaySuffix"`
				Speed         string `json:"speed"`
				Title         string `json:"title"`
			} `json:"speeds"`
			Intensity struct {
				Enumname    string `json:"enumName"`
				Description string `json:"description"`
			} `json:"intensity"`
			Nhqm string `json:"nhqm"`
		} `json:"devices"`
		Cpuminingenabled bool `json:"cpuMiningEnabled"`
		Cpuexists        bool `json:"cpuExists"`
		Stats            []struct {
			Statstime int64  `json:"statsTime"`
			Market    string `json:"market"`
			Algorithm struct {
				Enumname    string `json:"enumName"`
				Description string `json:"description"`
			} `json:"algorithm"`
			Unpaidamount             string  `json:"unpaidAmount"`
			Difficulty               float64 `json:"difficulty"`
			Proxyid                  int     `json:"proxyId"`
			Timeconnected            int64   `json:"timeConnected"`
			Xnsub                    bool    `json:"xnsub"`
			Speedaccepted            float64 `json:"speedAccepted"`
			Speedrejectedr1Target    float64 `json:"speedRejectedR1Target"`
			Speedrejectedr2Stale     float64 `json:"speedRejectedR2Stale"`
			Speedrejectedr3Duplicate float64 `json:"speedRejectedR3Duplicate"`
			Speedrejectedr4Ntime     float64 `json:"speedRejectedR4NTime"`
			Speedrejectedr5Other     float64 `json:"speedRejectedR5Other"`
			Speedrejectedtotal       float64 `json:"speedRejectedTotal"`
			Profitability            float64 `json:"profitability"`
		} `json:"stats"`
		Profitability      float64 `json:"profitability"`
		Localprofitability float64 `json:"localProfitability"`
		Rigpowermode       string  `json:"rigPowerMode"`
	} `json:"miningRigs"`
	Rignhmversions          []string `json:"rigNhmVersions"`
	Externaladdress         bool     `json:"externalAddress"`
	Totalprofitabilitylocal float64  `json:"totalProfitabilityLocal"`
	Pagination              struct {
		Size           int `json:"size"`
		Page           int `json:"page"`
		Totalpagecount int `json:"totalPageCount"`
	} `json:"pagination"`
}

func (m *MiningStats) ComputeTemperature() {
	for i := range m.Miningrigs {
		for y := range m.Miningrigs[i].Devices {
			if !(&Temperature{m.Miningrigs[i].Devices[y].Temperature}).IsRealTemp() {
				m.Miningrigs[i].Devices[y].GPUTemperature = int(m.Miningrigs[i].Devices[y].Temperature) & 0xffff
				m.Miningrigs[i].Devices[y].VRAMTemperature = int(m.Miningrigs[i].Devices[y].Temperature) >> 16
			}
		}
	}
}

func (t *Temperature) IsRealTemp() bool {
	return t.T < 1000000.0
}
