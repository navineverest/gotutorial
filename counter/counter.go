package counter

import "sync"

type Counter struct {
	mu         sync.Mutex
	currentVal int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.currentVal++
}

func (c *Counter) Value() int {
	return c.currentVal
}
