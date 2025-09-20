package utils

import (
	"fmt"
	"sync"
)

var mx sync.Mutex

func ThreadSavePrintln(message string) {
	mx.Lock()
	fmt.Println(message)
	mx.Unlock()
}

func ThreadSavePrintf(message string, args ...interface{}) {
	mx.Lock()
	fmt.Printf(message, args...)
	mx.Unlock()
}
