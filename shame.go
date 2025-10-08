package gg

//"fmt"

type Cache struct {
	data map[string]int
}

func New() *Cache {
	return &Cache{
		data: make(map[string]int),
	}
}

func (c *Cache) Set(key string, value int) {
	c.data[key] = value
}

func (c *Cache) Get(key string) int {
	val := c.data[key]
	return val
}

// удаление по ключу
func (c *Cache) Delete(key string) {
	delete(c.data, key)
}

// получение всей карты
func (c *Cache) All() map[string]int {
	return c.data
}
