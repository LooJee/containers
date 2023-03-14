package locker

type Locker interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}

var _ Locker = (*NoobLocker)(nil)

type NoobLocker struct {
}

func (l *NoobLocker) Lock()    {}
func (l *NoobLocker) Unlock()  {}
func (l *NoobLocker) RLock()   {}
func (l *NoobLocker) RUnlock() {}
