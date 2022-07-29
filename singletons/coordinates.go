package singleton

import "sync"

var (
	c    *Coordinates
	once sync.Once
)

func GetInstance() *Coordinates {
	once.Do(func() {
		c = &Coordinates{}
	})

	return c
}

type Coordinates struct {
	Satellites []Satellite
}

func (c *Coordinates) SetSatellites(satellites []Satellite) {
	c.Satellites = satellites
}

func (c *Coordinates) GetSatellites() (satellites []Satellite) {
	return c.Satellites
}
