package base

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuy(t *testing.T) {
	var wg sync.WaitGroup
	total := 1000
	goods := &Goods{}
	goods.InitializeTotal(100)
	for i := 0; i < total; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			go goods.Buy()
		}()
	}
	wg.Wait()
	assert.Equal(t, goods.total, 0)
}
