package unid

import (
	"sync/atomic"
	"time"
)

var suffix uint64 = 0

func rollSuffix() uint64 {
	atomic.AddUint64(&suffix, 1)
	atomic.CompareAndSwapUint64(&suffix, 100, 0)
	return suffix % 100
}

func Unid() uint64 {
	return uint64(time.Now().UnixNano())*100 + rollSuffix()
}
