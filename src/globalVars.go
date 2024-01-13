package main

import "time"

const (
	barbaraIDPath = "/etc/barbara_id.json"
	appConfigPath = "/appconfig/appconfig.json"
	logTag        = "logApp"
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
