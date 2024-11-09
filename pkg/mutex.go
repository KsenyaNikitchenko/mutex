package mutex

type Mutex struct {
	Count int
	done  chan bool
}

func NewMutex(count int) *Mutex {
	return &Mutex{
		Count: count,
		done:  make(chan bool, count),
	}
}

func (m *Mutex) Unlock() {
	m.done <- true
}

func (m *Mutex) Wait() {
	for i := 0; i < m.Count; i++ {
		<-m.done
	}
}
