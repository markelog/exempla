// Package ttl has some utils for working with time to live
package ttl

import "time"

// Get gets time to live
func Get(ttl int) time.Duration {
	return time.Now().Add(time.Millisecond * time.Duration(ttl)).Sub(time.Now())
}
