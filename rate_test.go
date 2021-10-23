package rate

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	r := New(100*time.Millisecond, 1*time.Second)
	assert.NotNil(t, r)
	assert.Equal(t, 10, len(r.samples))
}

func TestRate_Increment(t *testing.T) {
	interval := 1 * time.Second
	r := New(interval, 10*time.Second)
	stop := false
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for !stop {
			r.Increment()
		}
		wg.Done()
	}()
	time.AfterFunc(5*time.Second, func() {
		stop = true
	})
	wg.Wait()
	value := r.Value()
	fmt.Printf("rate %.2f/%v\n", value, interval)
	assert.Greater(t, int(value), 1000)
}

func TestRate_Value(t *testing.T) {
	interval, r := simulateRate()
	value := r.Value()
	fmt.Printf("rate %.2f/%v\n", value, interval)
	assert.InDelta(t, 100, value, 5)
}

func TestRate_ValueBy(t *testing.T) {
	_, r := simulateRate()
	value := r.ValueBy(time.Microsecond)
	fmt.Printf("rate %.4f/%v\n", value, time.Microsecond)
	assert.InDelta(t, 100.0/1000000.0, value, 0.001)
}

func TestRate_Value_AfterNoActivity(t *testing.T) {
	interval, r := simulateRate()
	time.Sleep(1 * time.Second)
	value := r.Value()
	fmt.Printf("rate %.2f/%v\n", value, interval)
	assert.Equal(t, 0, int(value))
}

func simulateRate() (time.Duration, *Rate) {
	interval := 100 * time.Millisecond
	r := New(interval, 1*time.Second)
	stop := false
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		var c int
		for !stop {
			r.Increment()
			c++
			if c > 100 {
				time.Sleep(100 * time.Millisecond)
				c = 0
			}
		}
		wg.Done()
	}()
	time.AfterFunc(1*time.Second, func() {
		stop = true
	})
	wg.Wait()
	return interval, r
}

func BenchmarkRate_Increment_100ms_1s(b *testing.B) {
	r := New(100*time.Millisecond, 1*time.Second)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Increment()
	}
}

func BenchmarkRate_Increment_1ms_1s(b *testing.B) {
	r := New(1*time.Millisecond, 1*time.Second)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Increment()
	}
}
