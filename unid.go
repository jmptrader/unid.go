// Copyright (c) 2014 by Kris Kovalik.

/*
Unid is a small and efficient package that generates timestamp based
identifiers with extra uniqueness ensure. The algorithm is pretty simple:
generate a unix nano timestamp, multiply it by 100 and add sequential,
atomically generated number from range 0-99. This way we have time based
and highly unique identfier in our hands.
*/
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

// Unid generates unix nano timestamp based id with ensured uniqueness.
// Returns 64-bit unsigned number.
func Unid() uint64 {
	return uint64(time.Now().UnixNano())*100 + rollSuffix()
}
