package gofillslicewithrand

import (
	"math/rand"
	"sync"
	"time"
)

var (
	fallbackLck  sync.Mutex
	fallbackRand *rand.Rand
)

func mathRandRead(p []byte) {
	fallbackLck.Lock()
	defer fallbackLck.Unlock()
	if fallbackRand == nil {
		fallbackRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	fallbackRand.Read(p)
}
