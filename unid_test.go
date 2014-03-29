package unid

import (
	"sync"
	"testing"
)

func TestSuffixRolling(t *testing.T) {
	for i := 0; i < 1001; i++ {
		suf := rollSuffix()
		if suf > 99 {
			t.Errorf("expected to get proper suffix")
		}
	}
	if suffix != 1 {
		t.Errorf("expected proper suffix rolling, got wrong value: %d", suffix)
	}
}

func TestSuffixAtomicity(t *testing.T) {
	var wg sync.WaitGroup
	n := 50000
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			rollSuffix()
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestUnid(t *testing.T) {
	u1 := Unid()
	if u1 == 0 {
		t.Errorf("expected to get propert unid")
	}
	u2 := Unid()
	if u2 == u1 {
		t.Errorf("expected to get unique unids")
	}
}

func TestUniquenessOfUnid(t *testing.T) {
	var wg sync.WaitGroup
	var m sync.Mutex
	n := 100000
	u := []uint64{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			id := Unid()
			m.Lock()
			u = append(u, id)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	for i := 0; i < len(u); i++ {
		for j := i + 1; j < len(u); j++ {
			if u[j] == u[i] {
				t.Errorf("expected uniq unids")
			}
		}
	}
}

func BenchmarkUnidGeneration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Unid()
	}
}
