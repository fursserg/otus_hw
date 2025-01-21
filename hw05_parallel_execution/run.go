package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type errorsCounter struct {
	ErrLimit int

	errCount int
	exceeded atomic.Bool
	mx       sync.Mutex
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup
	var result error
	ch := make(chan struct{}, n)
	ec := errorsCounter{ErrLimit: m}

	defer func() {
		close(ch)
	}()

	for _, v := range tasks {
		if ec.isErrorsExceeded() {
			break
		}

		ch <- struct{}{}
		wg.Add(1)

		go func() {
			defer func() {
				<-ch
				wg.Done()
			}()

			if ec.isErrorsExceeded() {
				return
			}

			if err := v(); err != nil {
				ec.increaseErrors()
			}
		}()
	}

	wg.Wait()

	if ec.isErrorsExceeded() {
		result = ErrErrorsLimitExceeded
	}

	return result
}

func (ec *errorsCounter) increaseErrors() {
	defer ec.mx.Unlock()
	ec.mx.Lock()

	ec.errCount++

	if ec.ErrLimit == ec.errCount {
		ec.exceeded.Store(true)
	}
}

func (ec *errorsCounter) isErrorsExceeded() bool {
	return ec.exceeded.Load()
}
