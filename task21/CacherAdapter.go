package main

type CacherAdapter struct {
	cacher *Cacher
}

func NewCacherAdapter(cacher *Cacher) *CacherAdapter {
	return &CacherAdapter{cacher: cacher}
}

func (c *CacherAdapter) Save(d *Data) {
	c.cacher.Cache(d)
}

func (c *CacherAdapter) Get(name string) {
	c.cacher.GetCache(name)
}
