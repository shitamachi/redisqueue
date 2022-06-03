package redisqueue

import (
	"os"
	"testing"
	"time"
)

func TestNewSignalHandler(t *testing.T) {
	t.Run("closes the returned channel on SIGINT", func(tt *testing.T) {
		ch := newSignalHandler()

		if p, err := os.FindProcess(os.Getpid()); err != nil {
			t.Fatal("Unable to find current process.", "pid", os.Getpid(), err)
		} else {
			// WARN: Sending Interrupt on Windows is not implemented. It will case the error of "timed out waiting for signal"
			_ = p.Signal(os.Interrupt)
		}
		select {
		case <-time.After(2 * time.Second):
			t.Error("timed out waiting for signal")
		case <-ch:
		}
	})
}
