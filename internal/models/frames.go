package models

import (
	"time"
)

type Frames struct {
	count     uint
	lastTime  int64
	perSecond uint
}

func (f *Frames) calculate() {
	now := time.Now().Unix()

	f.count++
	if now-f.lastTime >= 1 {
		f.perSecond = f.count
		f.count = 0
		f.lastTime = now
	}
}
