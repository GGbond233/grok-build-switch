//go:build windows

package singleinstance

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestAcquirePreventsDuplicateInstance(t *testing.T) {
	key := fmt.Sprintf("%s-%d-%d", t.Name(), os.Getpid(), time.Now().UnixNano())
	first, already, err := Acquire(key)
	if err != nil {
		t.Fatal(err)
	}
	if already {
		t.Fatal("first Acquire() reported an existing instance")
	}
	defer first.Close()

	second, already, err := Acquire(key)
	if err != nil {
		t.Fatal(err)
	}
	if second != nil {
		_ = second.Close()
	}
	if !already {
		t.Fatal("second Acquire() did not detect the existing instance")
	}
}
