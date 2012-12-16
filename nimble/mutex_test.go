package nimble

import (
	"math/rand"
	"testing"
	"time"
)

const D = 100 * time.Microsecond
const P = 0.95

// Write consecutive numbers in the array
// and test if read receives them in the correct order.
func TestRWMutex(t *testing.T) {
	N := 3
	a := make([]int, N)
	frames := 100
	m := newRWMutex(N)
	r1 := m.MakeRMutex()
	r2 := m.MakeRMutex()

	go write(m, a, N, frames)
	go read(r2, a, N, frames, t)
	read(r1, a, N, frames, t)
}

func write(m *rwMutex, a []int, N, frames int) {
	count := 0
	for i := 0; i < frames; i++ {
		prev := 0
		for j := 1; j <= N; j++ {
			m.next(1) // write
			//fmt.Printf("W % 3d % 3d: %d\n", prev, j, count)
			a[prev] = count
			m.done() // write
			count++
			prev = j
			if rand.Float32() > P {
				time.Sleep(D)
			}
		}
		if rand.Float32() > P {
			time.Sleep(D)
		}
	}
}

func read(m *rMutex, a []int, N, frames int, t *testing.T) {
	count := 0
	for i := 0; i < frames; i++ {
		prev := 0
		for j := 1; j <= N; j += 1 {
			m.next(1) // read
			//fmt.Printf("                   R % 3d % 3d: %d\n", prev, j, a[prev])
			if count != a[prev] {
				t.Error("got", a[prev], "expected", count)
			}
			m.done() // read
			count++
			prev = j
			if rand.Float32() > P {
				time.Sleep(D)
			}
		}
		if rand.Float32() > P {
			time.Sleep(D)
		}
	}
}
