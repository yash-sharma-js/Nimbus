package cache
// Local cache
import (
	"sync"
	"time"
)

type Item struct {
	Response []byte
	Headers  map[string]string
	Expiry   time.Time
}

type Cache struct {
	data map[string]Item
	mu   sync.RWMutex
}

func New() *Cache {
	return &Cache{
		data: make(map[string]Item),
	}
}

func (c *Cache) Get(key string) (Item, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.data[key]
	if !ok || time.Now().After(item.Expiry) {
		return Item{}, false
	}
	return item, true
}

func (c *Cache) Set(key string, item Item) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = item
}
