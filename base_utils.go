package main

import(
	"time"
	"log"
)

func Delay(ms int64) {
	time.Sleep(time.Duration(ms)*time.Millisecond)
}

func Log(toLog ...interface{}) {
	log.Println(toLog)
}
