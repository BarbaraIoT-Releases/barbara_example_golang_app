package main

import (
	"errors"
	"strconv"
	"time"
)

func main() {
	// show app name and app version
	initMessage()

	for {
		// check app config in loop until a valid app config is detected
		checkAppConfigErr := checkAppConfig()
		for checkAppConfigErr != nil {
			addLog(logTag, "failed to set app config: "+checkAppConfigErr.Error(), 1)
			time.Sleep(appSleepTime)
		}

		// show app config values as logs
		addLog(logTag, "Bool var value: "+strconv.FormatBool(varBool), 2)
		addLog(logTag, "String var value: "+varString, 2)
		addLog(logTag, "Int var value: "+strconv.Itoa(varInt), 2)

		// sleep and loop
		time.Sleep(appSleepTime)
	}

}

func checkAppConfig() error {
	debugLevelIface, debugLevelIfaceErr := getAppConfig("debugLevel")
	if debugLevelIfaceErr != nil {
		return errors.New("debugLevel: " + debugLevelIfaceErr.Error())
	}

	debugLevel, debugLevelErr = interfaceToInt(debugLevelIface)
	if debugLevelErr != nil {
		return errors.New("debugLevel: " + debugLevelErr.Error())
	}

	boolKeyIface, boolKeyIfaceErr := getAppConfig("boolKey")
	if boolKeyIfaceErr != nil {
		return errors.New("boolKey: " + boolKeyIfaceErr.Error())
	}

	varBool, varBoolErr = interfaceToBool(boolKeyIface)
	if varBoolErr != nil {
		return errors.New("varBool: " + varBoolErr.Error())
	}

	stringKeyIface, stringKeyIfaceErr := getAppConfig("stringKey")
	if stringKeyIfaceErr != nil {
		return errors.New("boolKey: " + stringKeyIfaceErr.Error())
	}

	varString, varStringErr = interfaceToString(stringKeyIface)
	if varStringErr != nil {
		return errors.New("varBool: " + varStringErr.Error())
	}

	intKeyIface, intKeyIfaceErr := getAppConfig("intKey")
	if intKeyIfaceErr != nil {
		return errors.New("intKey: " + intKeyIfaceErr.Error())
	}

	varInt, varIntErr = interfaceToInt(intKeyIface)
	if varIntErr != nil {
		return errors.New("varInt: " + varIntErr.Error())
	}

	return nil
}
