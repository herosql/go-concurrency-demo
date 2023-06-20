package base

import (
	"fmt"
	"sync"
)

type Goods struct {
	total int
	mutex sync.Mutex
}

func (g *Goods) InitializeTotal(initialTotal int) {
	g.total = initialTotal
}

func (g *Goods) Buy() {
	// 加锁
	g.mutex.Lock()
	// 解锁
	defer g.mutex.Unlock()
	if g.total > 0 {
		g.total = g.total - 1
	}
	fmt.Printf("当前库存%d\n", g.total)
}
