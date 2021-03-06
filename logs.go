package main

import "go.uber.org/zap"

//Check error with details informations
func Check(details string, err error) {
	if err != nil {
		zap.S().Errorf("Error %s %s", details, err.Error())
	} else {
		zap.S().Debugf("Success %s", details)
	}
}
