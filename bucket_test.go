package bucket

import (
	"testing"
	"time"
)

func TestBucket(t *testing.T) {
	bucket := New(10, time.Second, 100, true)
	if !bucket.Consume(100) {
		t.Fatal("init fill error")
	}
	time.Sleep(2 * time.Second)
	if !bucket.Consume(10) {
		t.Fatal("refill error")
	}
	if bucket.Consume(100) {
		t.Fatal("refill error")
	}
}
