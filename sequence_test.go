package sequence

import (
	"testing"
	"sync"
)

func TestSequence_GetSequence(t *testing.T) {
	seq1 := Int{}
	seq2 := Int{}

	if n := seq1.GetNext(); n != 1 {
		t.Errorf("Sequence1 starts from %d\n", n)
	}

	if n := seq2.GetNext(); n != 1 {
		t.Errorf("Sequence2 starts from %d\n", n)
	}

	for i := 1; i <= 10; i++ {
		if n := seq1.GetNext(); n != i + 1 {
			t.Error("Sequence1 next value invalid")
		}

		if n := seq2.GetNext(); n != i + 1 {
			t.Error("Sequence2 nenextt value invalid")
		}
	}
}

func TestSequence_GetSequenceParallelAccess(t *testing.T) {
	m := sync.Mutex{}
	n := []int{}

	appendSequence := func(j int) {
		m.Lock()
		n = append(n, j)
		m.Unlock()
	}

	c := 1000;
	s := Int{}
	t.Run("GetSequences", func(t *testing.T) {
		for i := 0; i < c; i++ {
			t.Run("GetSequenceParallel", func(t *testing.T) {
				t.Parallel()
				j := s.GetNext()
				appendSequence(j)
			})
		}
	})

	t.Run("CheckNextSequence", func(t *testing.T) {
		if j := s.GetNext(); j != c+1 {
			t.Errorf("GetNext invalid after 100 iterations. Expected: %d, actual: %d", c+1, j)
		}
	})
	t.Run("CheckSequences", func(t *testing.T) {
		if len(n) < c {
			t.Errorf("Int len invalid %d\n", len(n))
		}

		for i := 0; i < len(n); i++ {
			for j := 0; j < len(n); j++ {
				if i != j && n[i] == n[j] {
					t.Errorf("Equal sequence values in positions %d and %d\n", i, j)
				}
			}
		}
	})
}
