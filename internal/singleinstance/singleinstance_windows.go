//go:build windows

package singleinstance

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"golang.org/x/sys/windows"
)

type Lock struct {
	handle windows.Handle
}

// Acquire creates a per-session named mutex derived from key. alreadyRunning
// is true when another process still owns a mutex with the same name.
func Acquire(key string) (*Lock, bool, error) {
	sum := sha256.Sum256([]byte(key))
	name := `Local\grok_switch-` + hex.EncodeToString(sum[:12])
	namePtr, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return nil, false, err
	}
	handle, err := windows.CreateMutex(nil, true, namePtr)
	if errors.Is(err, windows.ERROR_ALREADY_EXISTS) {
		_ = windows.CloseHandle(handle)
		return nil, true, nil
	}
	if err != nil {
		return nil, false, err
	}
	return &Lock{handle: handle}, false, nil
}

func (l *Lock) Close() error {
	if l == nil || l.handle == 0 {
		return nil
	}
	handle := l.handle
	l.handle = 0
	releaseErr := windows.ReleaseMutex(handle)
	closeErr := windows.CloseHandle(handle)
	if releaseErr != nil {
		return releaseErr
	}
	return closeErr
}
