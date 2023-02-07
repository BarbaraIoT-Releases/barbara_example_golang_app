package main

import "time"

const (
	appConfigPath = "/appconfig/appconfig.json"
	logTag        = "baseApp"
	appSleepTime  = time.Duration(5) * time.Second
)

var (
	debugLevel    = 2
	debugLevelErr error
	varBool       bool
	varBoolErr    error
	varString     string
	varStringErr  error
	varInt        int
	varIntErr     error
)
