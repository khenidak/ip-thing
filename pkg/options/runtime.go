package options

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Runtime struct {
	// signaled if the process received SIGTERM/SIGINT
	stopChan chan os.Signal

	// root context is signaled when the StopChan is
	ctx context.Context

	// cancels the root context
	cancel func()

	// signaled when all clean up is done.
	doneChan chan struct{}

	mux     sync.Mutex
	bExited bool
}

// StopChan returns the runtime's Stop channel. Stop channel are signaled/closed
// when a 1) an os signal is received 2) ExitNow() is called
func (r *Runtime) StopChan() chan os.Signal {
	return r.stopChan
}

// Context returns a context that will be canceled
// if process is signaled or ExitNow() is called
func (r *Runtime) Context() context.Context {
	return r.ctx
}

// DoneChan returns a channel that will beclosed when
// SetDone is called. This is used by root func (aka main) to wait for
// a clean exit from all the go routine that may have been started
// if you use DoneChan you must call SetDone() in your main logic to signal
// those who are waiting
func (r *Runtime) DoneChan() chan struct{} {
	return r.doneChan
}

// ExitNow signals the process that a catastrophe has occured
// and it must exit all go routine that may have started
func (r *Runtime) ExitNow() chan struct{} {
	// allowing ExitNow to be called multiple times. Means
	// that caller can fail in multiple places at the same time
	r.mux.Lock()
	defer r.mux.Unlock()

	if r.bExited {
		return r.DoneChan()
	}

	r.cancel()
	close(r.stopChan) // signal stop

	r.bExited = true
	return r.DoneChan()
}

// SetDone signals all those who might be waiting for clean exit
// that exit has been completed
func (r *Runtime) SetDone() {
	close(r.doneChan)
}

func InitRuntime(r *Runtime) {
	// wire up signals
	r.stopChan = make(chan os.Signal, 1)
	r.doneChan = make(chan struct{})

	//
	bgContext := context.Background()
	r.ctx, r.cancel = signal.NotifyContext(bgContext, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(r.stopChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		select {
		case <-r.stopChan:
			r.ExitNow()
		}
	}()
}
