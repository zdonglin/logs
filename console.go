package logs

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type consoleWrite struct {
	sync.Mutex
	colorful bool
}

func newConsole() logger {
	cw := &consoleWrite{
		colorful: runtime.GOOS != "windows",
	}
	return cw
}

func (c *consoleWrite) init() error {
	if runtime.GOOS == "windows" {
		c.colorful = false
	} else {
		c.colorful = true
	}
	return nil
}

func (c *consoleWrite) writeMsg(when time.Time, msg string, level int) error {
	c.Lock()
	defer c.Unlock()
	if c.colorful {

	}
	b, _ := formatTimeHeader(when)
	fmt.Println(string(b), msg)
	return nil
}
func (c *consoleWrite) destroy() {

}
func (c *consoleWrite) flush() {

}
