package socketconnection

import "sync"

type SpawnController struct {
	Mutex       sync.Mutex
	shouldSpawn bool
}

func (s *SpawnController) GetShouldSpawn() bool {
	s.Mutex.Lock()
	result := s.shouldSpawn
	s.Mutex.Unlock()
	return result
}

func (s *SpawnController) SetValue(value bool) {
	s.Mutex.Lock()
	s.shouldSpawn = value
	s.Mutex.Unlock()
}
