// Package syncop simplify error handle for the case
// where error could raised from more than 1 go routine
package syncop

import "sync"

// Engine struct act as function receiver
type Engine struct {
	errChan chan error
	wg      *sync.WaitGroup
}

// New function create a new copy of engine
func New() Engine {

	return Engine{
		errChan: make(chan error),
		wg:      &sync.WaitGroup{},
	}

}

// ListenForError function will handle wait for wg
// and listen for any error that happen on go routines
// this function should called after last go routine declared
// and you can check the error from this function
func (e *Engine) ListenForError() error {

	go func(){
		e.wg.Wait()
		close(e.errChan)
	}()

	for err := range e.errChan {
		if err != nil {
			return err
		}
	}
	return nil
}

// HandleError function will send the error into channel that
// listened by ListenForError function
func (e *Engine) HandleError(err error) {
	e.errChan <- err
}

// WgAdd function wrap sync.WaitGroup's Add(delta int)
// this function should called before you spawnn new go routine
// so that go routine will registered as routine that will waited
func (e *Engine) WgAdd(delta int) {
	e.wg.Add(delta)
}

// WgDone function wrap sync.Waitgroup's Done()
// this function should called after your process done inside go routine
func (e *Engine) WgDone() {
	e.wg.Done()
}