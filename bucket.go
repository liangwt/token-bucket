package bucket

import (
	"time"
)

// Bucket 令牌桶结构
type Bucket struct {
	rate     int
	capacity int
	c        chan bool
}

// New 新建桶
func New(rate int, timeUnit time.Duration, capacity int, isFill bool) *Bucket {
	tokenChan := make(chan bool, capacity)

	b := &Bucket{
		rate:     rate,
		capacity: capacity,
		c:        tokenChan,
	}

	if isFill {
		for {
			select {
			case b.c <- true:
				continue
			default:
			}
			break
		}
	}

	go func() {
		ticker := time.NewTicker(1 * timeUnit / time.Duration(rate))
		for range ticker.C {
			select {
			case b.c <- true:
				continue
			default:
				continue
			}
		}
	}()

	return b
}

// Consume 消耗令牌
func (b *Bucket) Consume(size int) bool {
	if len(b.c) < size {
		return false
	}

	for i := size; i > 0; i-- {
		select {
		case <-b.c:
			continue
		default:
			return false
		}
	}

	return true
}
