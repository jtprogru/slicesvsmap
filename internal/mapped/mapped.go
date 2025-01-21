package mapped

type CacheMapped struct {
	IDs map[string]struct{}
}

func New() *CacheMapped {
	return &CacheMapped{
		IDs: make(map[string]struct{}, 100000000),
	}
}

func (c *CacheMapped) Add(key string) {
	if !c.Contains(key) {
		c.IDs[key] = struct{}{}
		return
	}
	return
}

func (c *CacheMapped) Contains(uid string) bool {
	_, ok := c.IDs[uid]
	return ok
}
