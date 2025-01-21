package sliced

import "slices"

type CacheSliced struct {
	IDs []string
}

func New() *CacheSliced {
	return &CacheSliced{
		IDs: make([]string, 0, 100000000),
	}
}

func (c *CacheSliced) Add(data string) {
	if !c.Contains(data) {
		c.IDs = append(c.IDs, data)
		slices.Sort(c.IDs)
		return
	}
	return
}

func (c *CacheSliced) Contains(uid string) bool {
	return slices.Contains(c.IDs, uid)
}
