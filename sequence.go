package sequence

import "sync"

type Int struct {
	sync.Mutex
	sequence int
}

func (s *Int) GetNext() int {
	s.Lock()
	defer s.Unlock()

	s.sequence++
	return s.sequence
}
