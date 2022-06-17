package pool

import (
	"math/rand"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	rand.Seed(time.Now().Unix())
	const totalCount = 20
	pool := NewPool(totalCount)

	for i := 0; i < totalCount; i++ {

		index, err := pool.GetIndex()
		if err != nil {
			t.Error(err)
		}
		t.Logf("index:%d", index)
	}
}
