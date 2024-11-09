package mutex

type Mutex struct {
	Count int
	done  chan bool
}

func (m *Mutex) Unlock() {
	m.done <- true
}

func (m *Mutex) Wait() {
	for i := 0; i < m.Count; i++ {
		<-m.done
	}
}
