package utils

import "log"

//CheckError 检查错误
func CheckError(err error){
	if err != nil {
		log.Panic(err)
	}
}