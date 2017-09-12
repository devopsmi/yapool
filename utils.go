package yapool

import "sync"

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrapper(fn func()) {
	w.Add(1)
	go func() {
		fn()
		w.Done()
	}()

}
