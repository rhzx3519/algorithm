package lock_free_queue

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"sync"
	"sync/atomic"
	"testing"
)

var dummyTask = func() error {
	return errors.New("dummy error")
}

func TestLockFreeQueue_Enqueue(t *testing.T) {
	que := NewLockFreeQueue()
	que.Enqueue(dummyTask)
	assert.False(t, que.Empty())
	assert.Equal(t, 1, que.Size())
}

func TestLockFreeQueue_multi_goroutine(t *testing.T) {
	que := NewLockFreeQueue()

	var wg sync.WaitGroup
	var N = 1000000
	var count int32
	countTask := func() error {
		atomic.AddInt32(&count, 1)
		return nil
	}

	for i := 0; i < N; i++ {
		wg.Add(1)
		go func() {
			que.Enqueue(countTask)
			wg.Done()
		}()
	}
	wg.Wait()

	assert.Equal(t, N, que.Size())
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func() {
			task := que.Dequeue()
			task()
			wg.Done()
		}()
	}
	wg.Wait()
	assert.True(t, que.Empty())

	assert.Equal(t, int32(N), count)
}
