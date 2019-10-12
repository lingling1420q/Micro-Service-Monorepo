package utils

import "monaco/logger"

/*ErrorCheck check if there is an error  */
func ErrorCheck(errMsg string, err error) {
	if err != nil {
		logger.Error(errMsg, err.Error())
		panic(err.Error())
	}
}
